package service

import (
	"book-service/src/common/errors"
	"book-service/src/common/helper"
	v "book-service/src/common/validator"
	"book-service/src/interface/repository"
	"book-service/src/interface/service"
	"book-service/src/model/dto"
	"book-service/src/model/entity"
	"context"

	pb "github.com/dwprz/synapsis-book-management-protobuf/protogen/book"
)

type BookImpl struct {
	bookRepo            repository.Book
	bookAggr            repository.BookAggregation
	popularTitleKeyRepo repository.PopularTitleKey
}

func NewBook(br repository.Book, bar repository.BookAggregation, pr repository.PopularTitleKey) service.Book {
	return &BookImpl{
		bookRepo:            br,
		bookAggr:            bar,
		popularTitleKeyRepo: pr,
	}
}

func (s *BookImpl) Add(ctx context.Context, data *dto.AddBookReq) error {
	if err := v.Validate.Struct(data); err != nil {
		return err
	}

	book, err := s.bookRepo.FindByTitle(ctx, data.Title)
	if err != nil {
		return err
	}

	if book != nil {
		return &errors.Response{HttpCode: 409, Message: "book already exists"}
	}

	err = s.bookRepo.Create(ctx, data)
	return err
}

func (s *BookImpl) FindMany(ctx context.Context, data *dto.GetBookReq) (*dto.DataWithPaging[[]*entity.Book], error) {
	if err := v.Validate.Struct(data); err != nil {
		return nil, err
	}

	limit, offset := helper.CreateLimitAndOffset(data.Page)

	if data.Title != "" {
		totalBooks, err := s.bookAggr.CountByTitle(ctx, data.Title)
		if err != nil {
			return nil, err
		}

		if totalBooks == 0 {
			return nil, &errors.Response{HttpCode: 404, Message: "books not found"}
		}

		books, err := s.bookRepo.FindManyByTitle(ctx, data.Title, limit, offset)
		if err != nil {
			return nil, err
		}

		go s.popularTitleKeyRepo.Upsert(context.Background(), data.Title)

		return helper.FormatPagedData(books, totalBooks, data.Page, limit), nil
	}

	if data.BookId != 0 || data.Author != "" || data.ISBN != "" || data.Location != "" || data.PublishedYear != 0 || data.Stock != nil {
		totalBooks, err := s.bookAggr.CountByFields(ctx, data)
		if err != nil {
			return nil, err
		}

		if totalBooks == 0 {
			return nil, &errors.Response{HttpCode: 404, Message: "books not found"}
		}

		books, err := s.bookRepo.FindManyByFields(ctx, data, limit, offset)
		if err != nil {
			return nil, err
		}

		return helper.FormatPagedData(books, totalBooks, data.Page, limit), nil
	}

	totalBooks, err := s.bookAggr.CountAll(ctx)
	if err != nil {
		return nil, err
	}

	if totalBooks == 0 {
		return nil, &errors.Response{HttpCode: 404, Message: "books not found"}
	}

	books, err := s.bookRepo.FindManyByRandom(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return helper.FormatPagedData(books, totalBooks, data.Page, limit), nil
}

func (s *BookImpl) FindManyByIds(ctx context.Context, data *dto.FindManyByIdsReq) ([]*pb.Book, error) {
	if err := v.Validate.Struct(data); err != nil {
		return nil, err
	}

	books, err := s.bookRepo.FindManyByIds(ctx, data.BookIds)
	return books, err
}

func (s *BookImpl) FindManyPopularBook(ctx context.Context) ([]*entity.Book, error) {
	books, err := s.popularTitleKeyRepo.FindManyPopularBook(ctx)

	return books, err
}

func (s *BookImpl) Update(ctx context.Context, data *dto.UpdateBookReq) (*entity.Book, error) {
	if err := v.Validate.Struct(data); err != nil {
		return nil, err
	}

	book, err := s.bookRepo.UpdateById(ctx, data)
	return book, err
}

func (s *BookImpl) Delete(ctx context.Context, bookId int) error {
	if err := v.Validate.Struct(&entity.BookId{Id: bookId}); err != nil {
		return err
	}

	err := s.bookRepo.DeleteById(ctx, bookId)
	return err
}
