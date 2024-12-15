package main

import "github.com/julienschmidt/httprouter"

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc("GET", "/api/v1/healthcheck", app.healthcheck)
	router.HandlerFunc("POST", "/api/v1/signup", app.signupHandler)
	// router.HandlerFunc("GET", "/api/v1/signin", app.signinHandler)
	// router.HandlerFunc("GET", "/api/v1/avatars", app.getAvailableAvatarsHandler)
	// router.HandlerFunc("GET", "/api/v1/elements", app.getAvailableElementsHandler)
	// admin routes
	// router.HandlerFunc("POST", "api/v1/admin/element", app.createElementHandler)
	// router.HandlerFunc("PUT", "api/v1/admin/element/:elementId", app.updateElementHandler)
	// router.HandlerFunc("POST", "api/v1/admin/avatar", app.createAvatarHandler)
	// router.HandlerFunc("POST", "api/v1/admin/map", app.createMapHandler)
	// user routes
	// router.HandlerFunc("POST", "api/v1/user/metadata", app.updateUserMetadataHandler)
	// router.HandlerFunc("GET", "api/v1/user/metadata/bulk", app.getUserMetadataBulkHandler)

	// space routes
	//router.HandlerFunc("POST", "api/v1/space/", app.createSpaceHandler)
	//router.HandlerFunc("DELETE", "api/v1/space/:spaceId", app.deleteSpaceHandler)
	//router.HandlerFunc("GET", "api/v1/space/all", app.getAvailableSpaceHandler)
	//router.HandlerFunc("GET", "api/v1/space/:spaceId", app.getSpaceByIdHandler)
	//router.HandlerFunc("POST", "api/v1/space/element", app.createSpaceElementHandler)
	//router.HandlerFunc("DELETE", "api/v1/space/element", app.deleteElementHandler)
	//
	//
	return router
}
