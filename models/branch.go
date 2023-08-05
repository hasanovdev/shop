package models

type CreateBranch struct {
	Name      string
	Address   string
	FoundedAt string
}

type Branch struct {
	Id        int
	Name      string
	Address   string
	CreatedAt string
	FoundedAt string
	Year      int
}

type IdRequest struct {
	Id int
}

type GetAllBranchRequest struct {
	Page  int
	Limit int
	Name  string
}

type GetAllBranch struct {
	Branches []Branch
	Count    int
}
