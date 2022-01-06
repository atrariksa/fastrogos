package handlers

import (
	"html/template"
	"net/http"

	"github.com/atrariksa/fastrogos/fenuma/configs"
	"github.com/sirupsen/logrus"
)

const (
	DASHBOARD_PAGE_TITLE string = "Dashboard"
	DASHBOARD_HTML       string = "dashboard.html"
)

type DashboardPage struct {
	Title      string
	UserMenu   template.HTML
	PluginMenu template.HTML
}

type DashboardPageHandler struct {
	cfg *configs.Config
	t   *template.Template
	log *logrus.Logger
}

func NewDashboardPageHandler(cfg *configs.Config, t *template.Template, log *logrus.Logger) *DashboardPageHandler {
	return &DashboardPageHandler{
		cfg: cfg,
		t:   t,
		log: log,
	}
}

func (lp *DashboardPageHandler) Handle(w http.ResponseWriter, r *http.Request) {
	dashboardPage := DashboardPage{Title: DASHBOARD_PAGE_TITLE}
	dashboardPage.UserMenu = `
	<li>
		<a href="#" id="id-manage-user">
			<span class="icon material-icons md-18">manage_accounts</span>
			<span class="title">User</span>
		</a>
    </li>
	`
	dashboardPage.PluginMenu = `
	<li>
		<a href="#">
			<span class="icon material-icons md-18">extension</span>
			<span class="title">Plugin</span>
		</a>
    </li>
	`
	err := lp.t.ExecuteTemplate(w, DASHBOARD_HTML, dashboardPage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}
