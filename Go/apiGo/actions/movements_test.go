package actions

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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
	DB := db.Test()
	handler := New(DB)
	id1 := uuid.New().String()
	id2 := uuid.New().String()
	id3 := uuid.New().String()
	id4 := uuid.New().String()
	rand.Seed(time.Now().UnixNano())
	team1 := create(id1, randomString(13), "club", "Colombia")
	team2 := create(id2, randomString(13), "club", "Colombia")

	team3 := create(id3, randomString(13), "national", "")
	team4 := create(id4, randomString(13), "national", "")

	player1 := createPlayer(id1)

	tests := []tests{

		{
			//CLub sign player
			input: models.PlayerTeam{
				PlayerID: player1.ID,
				TeamID:   team1.ID,
			},
			status: http.StatusOK,
		},
		{
			//National sign player
			input: models.PlayerTeam{
				PlayerID: player1.ID,
				TeamID:   team3.ID,
			},
			status: http.StatusOK,
		},

		{
			//Player is already in this team.
			input: models.PlayerTeam{
				PlayerID: player1.ID,
				TeamID:   team1.ID,
			},
			status: http.StatusBadRequest,
		},
		{
			//Player can't be in two clubs.
			input: models.PlayerTeam{
				PlayerID: player1.ID,
				TeamID:   team2.ID,
			},
			status: http.StatusBadRequest,
		},
		{
			//Player can't be in two national teams.
			input: models.PlayerTeam{
				PlayerID: player1.ID,
				TeamID:   team4.ID,
			},
			status: http.StatusBadRequest,
		},
		{
			//PlayerID cant not be empty.
			input: models.PlayerTeam{
				PlayerID: uuid.Nil,
				TeamID:   team1.ID,
			},
			status: http.StatusBadRequest,
		},
		{
			//TeamID cant not be empty.
			input: models.PlayerTeam{
				PlayerID: player1.ID,
				TeamID:   uuid.Nil,
			},
			status: http.StatusBadRequest,
		},
		{
			//Player does not exists.
			input: models.PlayerTeam{
				PlayerID: uuid.MustParse("7998808a-18c9-4e19-8954-41ca344e1276"),
				TeamID:   team1.ID,
			},
			status: http.StatusNotFound,
		},
		{
			//Team does not exists.
			input: models.PlayerTeam{
				PlayerID: player1.ID,
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

	id1 := uuid.New().String()
	//id2 := uuid.New().String()
	id3 := uuid.New().String()
	id4 := uuid.New().String()
	rand.Seed(time.Now().UnixNano())
	team1 := create(id1, randomString(13), "club", "Colombia")
	//team2 := create(id2, randomString(13), "club", "Colombia")

	team3 := create(id3, randomString(13), "national", "")
	team4 := create(id4, randomString(13), "national", "")

	player1 := createPlayer(id1)

	tests := []tests{
		{
			//Player is not in this club.
			input: models.PlayerTeam{
				PlayerID: player1.ID,
				TeamID:   team3.ID,
			},
			status: http.StatusBadRequest,
		},
		{
			//Player does not exists.
			input: models.PlayerTeam{
				PlayerID: uuid.MustParse("0065522b-6946-483b-9f60-8d61c1e62478"),
				TeamID:   team1.ID,
			},
			status: http.StatusNotFound,
		},
		{
			//Player is not in this national.
			input: models.PlayerTeam{
				PlayerID: player1.ID,
				TeamID:   team4.ID,
			},
			status: http.StatusBadRequest,
		},
		{
			//PlayerID cant not be empty.
			input: models.PlayerTeam{
				PlayerID: uuid.Nil,
				TeamID:   team1.ID,
			},
			status: http.StatusBadRequest,
		},
		{
			//TeamID cant not be empty.
			input: models.PlayerTeam{
				PlayerID: player1.ID,
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
