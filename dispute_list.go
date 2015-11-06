package omise

/*
!!! DO NOT MODIFY !!!

autogenerated
 src: gen_list_job.tmpl
 job: &main.GenListJob{Name:"Dispute"}
  on: Fri Nov 06 13:48:24 +0700 2015
  by: chakrit
*/

// DisputeList represents the list structure returned by Omise's REST API that contains
// Dispute struct as member elements. See the pagination and lists documentation at
// https://www.omise.co/api-pagination for more information.
type DisputeList struct {
	List
	Data []*Dispute `json:"data"`
}

// Find finds and returns Dispute with the given id. Returns nil if not found.
func (list *DisputeList) Find(id string) *Dispute {
	for _, item := range list.Data {
		if item.ID == id {
			return item
		}
	}

	return nil
}
