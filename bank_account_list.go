package omise

/*
!!! DO NOT MODIFY !!!

autogenerated
 src: gen_list_job.tmpl
 job: &main.GenListJob{Name:"BankAccount"}
  on: Fri Nov 06 13:48:24 +0700 2015
  by: chakrit
*/

// BankAccountList represents the list structure returned by Omise's REST API that contains
// BankAccount struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type BankAccountList struct {
	List
	Data []*BankAccount `json:"data"`
}

// Find finds and returns BankAccount with the given id. Returns nil if not found.
func (list *BankAccountList) Find(id string) *BankAccount {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}
