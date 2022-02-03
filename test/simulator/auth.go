package simulator

import (
	"encoding/json"
	"net/http"

	v1 "go.reefassistant.com/apex/client/private/v1alpha1"
)

type noUserResponse struct {
	Username string `json:"username"`
}

// InvalidateSession kills the given session.
func (s *Controller) InvalidateSession(sid string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.sids, sid)
}

// InvalidateAllSessions kills all existing sessions.
func (s *Controller) InvalidateAllSessions() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.sids = make(map[string]struct{})
}

func (s *Controller) handlePostLogin(w http.ResponseWriter, r *http.Request) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Invalid JSON request just yields an empty response, no error code nor body.
	var loginRequest v1.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		return
	}

	// Successful authentication create a new 25 character session.
	if loginRequest.Login == s.login && loginRequest.Password == s.password {

		// Generate new sid and store it.
		sid := randSid()
		s.sids[sid] = struct{}{}

		// Attach the cookie
		http.SetCookie(w, &http.Cookie{
			Name:  COOKIE_SID,
			Value: sid,
		})

		// Sid JSON response
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(v1.LoginResponse{ConnectSid: sid})
		return
	}

	// Invalid credentials, unauthorized.
	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}

func (s *Controller) handleGetUser(w http.ResponseWriter, r *http.Request) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	w.Header().Set("content-type", "application/json")

	// If no cookie, return just empty username with 200 OK.
	cookie, err := r.Cookie(COOKIE_SID)
	if err != nil {
		json.NewEncoder(w).Encode(noUserResponse{})
		return
	}

	// For a valid session return the response.
	if _, ok := s.sids[cookie.Value]; ok {
		json.NewEncoder(w).Encode(v1.GetUserResponse{Username: s.login})
		return
	}

	// Invalid session still yields an empty username with 200 OK.
	json.NewEncoder(w).Encode(noUserResponse{})
}
