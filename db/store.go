package db

import "context"

type UserStore interface {
	GetUser(ctx context.Context, username string) (*User, error) 
	CreateUser(ctx context.Context, arg CreateUserParams) (*User, error)
	ListUser(ctx context.Context, arg ListUserParams) ([]*User, error)
}

type VoteStore interface {
	DoVoteTx(ctx context.Context, arg DoVoteTxParams) error
	CreateFormAndOptionTx(ctx context.Context, arg CreateFormAndOptionTxParams) error
	DeleteVoteTx(ctx context.Context, formId int64) error
	EndVoteTx(ctx context.Context) error


	ListForms(ctx context.Context) ([]*Form, error)
	GetForm(ctx context.Context, id int64) (*Form, error)
	GetOptionByFormId(ctx context.Context, formId int64) ([]*Option, error)
	GetUserVoteRecord(ctx context.Context, arg GetUserVoteRecordParams) ([]*VoteRecord, error)

	UpdateForm(ctx context.Context, arg UpdateFormParams) error
}

/*
	AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Account, error)
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateVerifyEmail(ctx context.Context, arg CreateVerifyEmailParams) (VerifyEmail, error)
	DeleteAccount(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (Account, error)
	GetAccountForUpdate(ctx context.Context, id int64) (Account, error)
	GetEntry(ctx context.Context, id int64) (Entry, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetTransfer(ctx context.Context, id int64) (Transfer, error)
	
	GetUser(ctx context.Context, username string) (User, error)
	
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateVerifyEmail(ctx context.Context, arg UpdateVerifyEmailParams) (VerifyEmail, error)

	*/