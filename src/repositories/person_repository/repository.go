package person_repository

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	rb_cache "rinha-de-backend/pkg/cache"
	"rinha-de-backend/pkg/database"
	"rinha-de-backend/src/dtos/person_dto"
	"rinha-de-backend/src/structs"
)

type Repository struct {
	cache *rb_cache.Redis
	DB    *database.Postgres
}

func New(DB *database.Postgres, cache *rb_cache.Redis) *Repository {
	return &Repository{
		DB:    DB,
		cache: cache,
	}
}

func (r Repository) GetByNickname(ctx context.Context, nickname string) (structs.Person, error) {
	var person structs.Person
	//err := r.cache.Cache.Get(ctx, nickname, &person)
	//if err != nil {
	//	if !errors.Is(cache.ErrCacheMiss, err) {
	//		log.Error(err)
	//		return structs.Person{}, err
	//	}
	//}

	db := r.DB.DB

	tx := db.Model(&structs.Person{}).
		Select("*").
		Where("nickname = ?", nickname).
		First(&person)
	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return structs.Person{}, tx.Error
	}

	return person, nil
}

func (r Repository) GetByID(ctx context.Context, ID uuid.UUID) (structs.Person, error) {
	var person structs.Person

	//err := r.cache.Cache.Get(ctx, ID.String(), &person)
	//if err != nil {
	//	if !errors.Is(cache.ErrCacheMiss, err) {
	//		return structs.Person{}, err
	//	}
	//}

	db := r.DB.DB
	tx := db.Model(&structs.Person{}).
		Select("*").
		Where("id = ?", ID.String()).
		First(&person)
	if tx.Error != nil {
		return structs.Person{}, tx.Error
	}

	if person.ID == uuid.Nil {
		return structs.Person{}, fmt.Errorf("not found")
	}

	return person, nil
}

func (r Repository) List(ctx context.Context, params person_dto.ListRequestParams) ([]structs.Person, error) {
	var persons []structs.Person

	db := r.DB.DB

	param := params.Params

	qb := db.Table("people").
		Select("*").
		Preload("Stack").
		Where("people.name LIKE ? OR people.nickname LIKE ?", "%"+param+"%", "%"+param+"%").
		Or("people.id IN (SELECT person_id FROM stacks WHERE LOWER(name) LIKE ?)", "%"+strings.ToLower(param)+"%")

	if params.Size != 0 {
		qb = qb.Limit(params.Size)
	}

	tx := qb.Find(&persons)

	if tx.Error != nil {
		return []structs.Person{}, tx.Error
	}

	return persons, nil
}

func (r Repository) Save(ctx context.Context, payload person_dto.CreatePayload) (structs.Person, error) {
	stacks := make([]structs.Stack, 0)
	for _, stack := range payload.Stack {
		stacks = append(stacks, structs.Stack{
			ID:   uuid.New(),
			Name: stack,
		})
	}

	birtDate, _ := time.Parse("2006-01-02", payload.BirthDate)

	ID := uuid.New()
	person := structs.Person{
		ID:        ID,
		Name:      payload.Name,
		Nickname:  payload.Nickname,
		BirthDate: birtDate,
		Stack:     stacks,
	}

	tx := r.DB.DB.Create(&person)
	if tx.Error != nil {
		return structs.Person{}, tx.Error
	}

	//err := r.cache.Cache.Set(&cache.Item{
	//	Ctx:   ctx,
	//	Key:   payload.Nickname,
	//	Value: person,
	//	TTL:   time.Hour,
	//})
	//
	//err = r.cache.Cache.Set(&cache.Item{
	//	Ctx:   ctx,
	//	Key:   ID.String(),
	//	Value: person,
	//	TTL:   time.Hour,
	//})

	//if err != nil {
	//	return structs.Person{}, err
	//}

	return person, nil
}
