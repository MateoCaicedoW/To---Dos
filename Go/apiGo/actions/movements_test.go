package actions

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mateo/apiGo/db"
	"github.com/mateo/apiGo/models"
)

type tests struct {
	input  models.PlayerTeam
	status int
}

func TestSignPlayer(T *testing.T) {
	DB := db.Init()
	handler := New(DB)

	tests := []tests{

		{
			//Player is already in this team.
			input: models.PlayerTeam{
				PlayerID: uuid.MustParse("834efe36-999e-42ea-9f17-6aec5d7b0612"),
				TeamID:   uuid.MustParse("aed3c9e6-4981-41f5-b587-57be70341744"),
			},
			status: http.StatusBadRequest,
		},
		{
			//Player can't be in two clubs.
			input: models.PlayerTeam{
				PlayerID: uuid.MustParse("0065522b-6946-483b-9f60-8d61c1e62459"),
				TeamID:   uuid.MustParse("7998808a-18c9-4e19-8954-41ca344e1276"),
			},
			status: http.StatusBadRequest,
		},
		{
			//Player can't be in two national teams.
			input: models.PlayerTeam{
				PlayerID: uuid.MustParse("d1b1c06b-37d9-4bdd-8dad-4d4320c5e178"),
				TeamID:   uuid.MustParse("76847802-0f35-4443-a061-b2adf706a431"),
			},
			status: http.StatusBadRequest,
		},
		{
			//PlayerID cant not be empty.
			input: models.PlayerTeam{
				PlayerID: uuid.Nil,
				TeamID:   uuid.MustParse("7998808a-18c9-4e19-8954-41ca344e1276"),
			},
			status: http.StatusBadRequest,
		},
		{
			//TeamID cant not be empty.
			input: models.PlayerTeam{
				PlayerID: uuid.MustParse("0065522b-6946-483b-9f60-8d61c1e62459"),
				TeamID:   uuid.Nil,
			},
			status: http.StatusBadRequest,
		},
		{
			//Player does not exists.
			input: models.PlayerTeam{
				PlayerID: uuid.MustParse("7998808a-18c9-4e19-8954-41ca344e1276"),
				TeamID:   uuid.MustParse("7998808a-18c9-4e19-8954-41ca344e1276"),
			},
			status: http.StatusNotFound,
		},
		{
			//Team does not exists.
			input: models.PlayerTeam{
				PlayerID: uuid.MustParse("0065522b-6946-483b-9f60-8d61c1e62459"),
				TeamID:   uuid.MustParse("0065522b-6946-483b-9f60-8d61c1e62459"),
			},
			status: http.StatusNotFound,
		},
	}
	for i, item := range tests {
		jsonStr, err := json.Marshal(item.input)
		if err != nil {
			panic(err)
		}
		router := mux.NewRouter()
		router.HandleFunc("/team/sign-player", handler.SignPlayer).Methods(http.MethodPost)

		server := &http.Server{
			Addr:    ":3000",
			Handler: router,
		}
		requestresponse := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/team/sign-player", bytes.NewBuffer(jsonStr))

		server.Handler.ServeHTTP(requestresponse, req)

		if status := requestresponse.Code; status != item.status {
			T.Errorf("Test %d handler returned wrong status code: got %v want %v", i, status, item.status)
		}
		log.Printf("Test %d: StatusCode: %v", i, requestresponse.Code)

	}

}

func TestUnsignPlayer(T *testing.T) {
	DB := db.Init()
	handler := New(DB)

	tests := []tests{
		{
			//Player is not in this team.
			input: models.PlayerTeam{
				PlayerID: uuid.MustParse("0065522b-6946-483b-9f60-8d61c1e62459"),
				TeamID:   uuid.MustParse("7998808a-18c9-4e19-8954-41ca344e1276"),
			},
			status: http.StatusBadRequest,
		},
		{
			//Player does not exists.
			input: models.PlayerTeam{
				PlayerID: uuid.MustParse("0065522b-6946-483b-9f60-8d61c1e62478"),
				TeamID:   uuid.MustParse("7998808a-18c9-4e19-8954-41ca344e1276"),
			},
			status: http.StatusNotFound,
		},
		{
			//Player is not in this team.
			input: models.PlayerTeam{
				PlayerID: uuid.MustParse("0065522b-6946-483b-9f60-8d61c1e62459"),
				TeamID:   uuid.MustParse("7998808a-18c9-4e19-8954-41ca344e1276"),
			},
			status: http.StatusBadRequest,
		},
		{
			//PlayerID cant not be empty.
			input: models.PlayerTeam{
				PlayerID: uuid.Nil,
				TeamID:   uuid.MustParse("7998808a-18c9-4e19-8954-41ca344e1276"),
			},
			status: http.StatusBadRequest,
		},
		{
			//TeamID cant not be empty.
			input: models.PlayerTeam{
				PlayerID: uuid.MustParse("0065522b-6946-483b-9f60-8d61c1e62459"),
				TeamID:   uuid.Nil,
			},
			status: http.StatusBadRequest,
		},
	}

	for i, item := range tests {
		jsonStr, err := json.Marshal(item.input)
		if err != nil {
			panic(err)
		}
		router := mux.NewRouter()
		router.HandleFunc("/team/unsign-player", handler.UnsignPlayer).Methods(http.MethodPost)

		server := &http.Server{
			Addr:    ":3000",
			Handler: router,
		}
		requestresponse := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/team/unsign-player", bytes.NewBuffer(jsonStr))

		server.Handler.ServeHTTP(requestresponse, req)

		if status := requestresponse.Code; status != item.status {
			T.Errorf("Test %d handler returned wrong status code: got %v want %v", i, status, item.status)
		}
		log.Printf("Test %d: StatusCode: %v", i, requestresponse.Code)
	}

}

func TestTransferPlayer(t *testing.T) {
	DB := db.Init()
	handler := New(DB)

	tests := []tests{
		{
			//player is in this team.
			input: models.PlayerTeam{
				PlayerID: uuid.MustParse("0065522b-6946-483b-9f60-8d61c1e62459"),
				TeamID:   uuid.MustParse("7998808a-18c9-4e19-8954-41ca344e1276"),
			},
			status: http.StatusBadRequest,
		},
		{
			//player does not exists.
			input: models.PlayerTeam{
				PlayerID: uuid.MustParse("0065522b-6946-483b-9f60-8d61c1e62478"),
				TeamID:   uuid.MustParse("7998808a-18c9-4e19-8954-41ca344e1276"),
			},
			status: http.StatusNotFound,
		},
	}

	for i, item := range tests {
		jsonStr, err := json.Marshal(item.input)
		if err != nil {
			panic(err)
		}
		router := mux.NewRouter()
		router.HandleFunc("/team/transfer-player", handler.TransferPlayer).Methods(http.MethodPost)

		server := &http.Server{
			Addr:    ":3000",
			Handler: router,
		}
		requestresponse := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/team/transfer-player", bytes.NewBuffer(jsonStr))

		server.Handler.ServeHTTP(requestresponse, req)

		if status := requestresponse.Code; status != item.status {
			t.Errorf("Test %d handler returned wrong status code: got %v want %v", i, status, item.status)
		}
		log.Printf("Test %d: StatusCode: %v", i, requestresponse.Code)
	}
}
