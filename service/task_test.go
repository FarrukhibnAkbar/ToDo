package service

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/FarrukhibnAkbar/ToDo/genproto"
)

func TestTaskService_Create(t *testing.T) {
	tests := []struct {
		result string
		input  pb.Task
		want   pb.Task
	}{
		{
			result: "successful",
			input: pb.Task{
				Assignee: "Farrux",
				Title:    "Gopher",
				Summary:  "Some ...",
				Deadline: "2021-12-20T00:00:00Z",
				Status:   "active",
			},
			want: pb.Task{
				Assignee:  "Farrux",
				Title:     "Gopher",
				Summary:   "Some ...",
				Deadline:  "2021-12-20T00:00:00Z",
				Status:    "active",
				CreatedAt: "2021-12-20",
				UpdatedAt: "2021-12-20",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.result, func(t *testing.T) {
			got, err := client.Create(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to create task", err)
			}
			got.Id = ""
			got.CreatedAt = "2021-12-20"
			got.UpdatedAt = "2021-12-20"
			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v, \ngot: %v", tc.result, tc.want, *got)
			}
		})
	}
}

func TestTaskService_Get(t *testing.T) {
	tests := []struct {
		result string
		input  pb.ByIdReq
		want   pb.Task
	}{
		{
			result: "successful",
			input:  pb.ByIdReq{Id: "1fdf46ed-901c-4b2a-8ec0-ef1782b9d47e"},
			want: pb.Task{
				Id:        "1fdf46ed-901c-4b2a-8ec0-ef1782b9d47e",
				Assignee:  "Assignee",
				Title:     "Title",
				Summary:   "Summary",
				Deadline:  "2021-12-20T00:00:00Z",
				Status:    "active",
				CreatedAt: "2021-12-21T11:24:17.933594Z",
				UpdatedAt: "0001-01-01 00:00:00 +0000 UTC",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.result, func(t *testing.T) {
			got, err := client.Get(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to Get task", err)
			}

			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expexted: %v, \ngot: %v", tc.result, tc.want, *got)
			}
		})
	}

}

func TestTaskService_List(t *testing.T) {
	tests := []struct {
		result string
		input  pb.ListReq
		want   []pb.Task
	}{
		{
			result: "successful",
			input: pb.ListReq{
				Page:  1,
				Limit: 3,
			},
			want: []pb.Task{
				{
					Id:        "1c07daf8-c59c-48a0-bbd1-55f2c2824293",
					Assignee:  "Farrux",
					Title:     "Gopher",
					Summary:   "Some ...",
					Deadline:  "2002-02-02T00:00:00Z",
					Status:    "active",
					CreatedAt: "2021-12-21T10:13:03.970075Z",
				},
				{
					Id:        "a9f827d2-0785-450f-a532-170dd0777657",
					Assignee:  "Farrux",
					Title:     "Gopher",
					Summary:   "Some ...",
					Deadline:  "2002-02-02T00:00:00Z",
					Status:    "active",
					CreatedAt: "2021-12-21T10:48:17.097422Z",
				},
				{
					Id:        "f3152ceb-adf4-4241-845f-44775e664a77",
					Assignee:  "Assignee",
					Title:     "Title",
					Summary:   "Summary",
					Deadline:  "2021-12-20T00:00:00Z",
					Status:    "active",
					CreatedAt: "2021-12-21T10:53:56.232525Z",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.result, func(t *testing.T) {
			got, err := client.List(context.Background(), &tc.input)
			for i := range got.Tasks {
				if err != nil {
					t.Error("failed to Get Lists of task", err)
				}
				if !reflect.DeepEqual(tc.want[i], *got.Tasks[i]) {
					t.Fatalf("%s: expected: %v, \ngot: %v", tc.result, tc.want[i], *got.Tasks[i])
				}
			}
		})
	}
}

func TestTaskService_Update(t *testing.T) {
	tests := []struct {
		result string
		input  pb.Task
		want   pb.Task
	}{
		{
			result: "successful",
			input: pb.Task{
				Id:       "23ebb512-05d4-4cdb-bf92-a674b2614e43",
				Assignee: "Farrux",
				Title:    "Gopher",
				Summary:  "Some ...",
				Deadline: "2002-02-02T00:00:00Z",
				Status:   "active",
			},
			want: pb.Task{
				Id:        "23ebb512-05d4-4cdb-bf92-a674b2614e43",
				Assignee:  "Farrux",
				Title:     "Gopher",
				Summary:   "Some ...",
				Deadline:  "2002-02-02T00:00:00Z",
				Status:    "active",
				CreatedAt: "2021-12-21T10:50:17.440705Z",
				UpdatedAt: "2021-12-21",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.result, func(t *testing.T) {
			got, err := client.Update(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to update task", err)
			}
			got.UpdatedAt = "2021-12-21"
			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v, \n got: %v", tc.result, tc.want, *got)
			}
		})
	}
}

func TestTaskService_Delete(t *testing.T) {
	tests := []struct {
		result string
		input  pb.ByIdReq
		want   pb.EmptyResp
	}{
		{
			result: "successful",
			input: pb.ByIdReq{
				Id: "528a1bb9-01e6-4bed-b85f-e521a893da54",
			},
			want: pb.EmptyResp{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.result, func(t *testing.T) {
			got, err := client.Delete(context.Background(), &tc.input)
			if err != nil {
				t.Error("failed to delete task", err)
			}
			if !reflect.DeepEqual(tc.want, *got) {
				t.Fatalf("%s: expected: %v, \n got: %v", tc.result, tc.want, *got)
			}
		})

	}
}

func TestTaskService_ListOverdue(t *testing.T) {
	tests := []struct {
		result string
		input  pb.ListOverReq
		want   []pb.Task
	}{
		{
			result: "successful",
			input:  pb.ListOverReq{Time: "2021-12-21", Page: 1, Limit: 2},
			want: []pb.Task{
				{
					Id:        "23ebb512-05d4-4cdb-bf92-a674b2614e43",
					Assignee:  "Farrux",
					Title:     "Gopher",
					Summary:   "Some ...",
					Deadline:  "2002-02-02T00:00:00Z",
					Status:    "active",
					CreatedAt: "2021-12-21T10:50:17.440705Z",
					UpdatedAt: "2021-12-21",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.result, func(t *testing.T) {
			got, err := client.ListOverdue(context.Background(), &tc.input)
			for i := range got.Tasks {
				if err != nil {
					t.Error("failed to Get Lists of task", err)
				}
				if !reflect.DeepEqual(tc.want[i], *got.Tasks[i]) {
					t.Fatalf("%s: expected: %v, \ngot: %v", tc.result, tc.want[i], *got.Tasks[i])
				}
			}
		})
	}
}
