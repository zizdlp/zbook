package db

// BinarySearch 使用二分查找在已排序的数组中查找指定值的插入位置
// 如果找到该值，则返回其索引；否则返回应插入的位置索引
// QueryMd5ForCheckRow:RelativePath 是升序
func BinarySearch(arr []QueryMd5ForCheckRow, length int, value string) int {
	left := 0
	right := length - 1
	mid := -1
	for left < right {
		mid = (left + right) / 2
		if value > arr[mid].RelativePath {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left // 返回应插入的位置索引
}

type CreateParams struct {
	RelativePath []string
	UserID       []int64
	RepoID       []int64
	MainContent  []string
	TableContent []string
	Md5          []string
	VersionKey   []string
}

func (params *CreateParams) Append(relativePath string, userID int64, repoID int64, mainContent string, tableContent string, md5 string, versionKey string) {
	params.RelativePath = append(params.RelativePath, relativePath)
	params.UserID = append(params.UserID, userID)
	params.RepoID = append(params.RepoID, repoID)
	params.MainContent = append(params.MainContent, mainContent)
	params.TableContent = append(params.TableContent, tableContent)
	params.Md5 = append(params.Md5, md5)
	params.VersionKey = append(params.VersionKey, versionKey)
}

type UpdateParams struct {
	RelativePath []string
	RepoID       []int64
	MainContent  []string
	TableContent []string
	Md5          []string
	VersionKey   []string
}

func (params *UpdateParams) Append(RelativePath string, RepoID int64, MainContent string, TableContent string, Md5 string, VersionKey string) {
	params.RelativePath = append(params.RelativePath, RelativePath)
	params.RepoID = append(params.RepoID, RepoID)
	params.MainContent = append(params.MainContent, MainContent)
	params.TableContent = append(params.TableContent, TableContent)
	params.Md5 = append(params.Md5, Md5)
	params.VersionKey = append(params.VersionKey, VersionKey)
}

type KeyParams struct {
	RelativePath []string
	RepoID       []int64
	VersionKey   []string
}

func (params *KeyParams) Append(relativePath string, repoID int64, versionKey string) {
	params.RelativePath = append(params.RelativePath, relativePath)
	params.RepoID = append(params.RepoID, repoID)
	params.VersionKey = append(params.VersionKey, versionKey)
}
