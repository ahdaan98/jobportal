package handler

import (
	"net/http"

	"github.com/ahdaan67/jobportal/internal/gateway/handler/employer"
	"github.com/ahdaan67/jobportal/internal/gateway/handler/job"
	"github.com/ahdaan67/jobportal/internal/gateway/handler/jobseeker"
	"github.com/ahdaan67/jobportal/internal/gateway/handler/newsletter"
	"github.com/ahdaan67/jobportal/internal/gateway/middleware"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func RegisterRoutes(jobhandler *job.Handler, jobseekerhandler *jobseeker.Handler, employerhandler *employer.Handler, videoHandler *VideoCallHandler, newletter *newsletter.Handler) *gin.Engine {
	r = gin.Default()

	r.Static("/static", "internal/gateway/static")
	r.LoadHTMLGlob("internal/gateway/template/*")

	r.GET("/", func(ctx *gin.Context) { ctx.HTML(http.StatusOK, "index.html", nil) })

	r.GET("/exit", videoHandler.ExitPage)
	r.GET("/error", videoHandler.ErrorPage)
	r.GET("/video", videoHandler.IndexedPage)

	r.POST("/jobseeker/signup", jobseekerhandler.JobSeekerSignup)
	r.POST("/jobseeker/verify/otp", jobseekerhandler.VerifyOTP)

	r.POST("/jobseeker/login", jobseekerhandler.LoginJobseeker)

	r.POST("/jobseeker/profile", jobseekerhandler.CreateJobSeekerProfile)

	r.GET("/linkedin/signin", jobseekerhandler.LinkedinSignIn)
	r.GET("/linkedin/callback", jobseekerhandler.LinkedInSignInCallback)

	r.GET("/jobseeker/newsletter/razorpay/:subscription", newletter.RazorpayPayment)
	r.GET("/jobseeker/newsletter/razorpay/verify/payment", newletter.VerifyPayment)

	r.GET("/job/:id", jobhandler.GetJob)
	r.GET("/jobs", jobhandler.ListJobs)

	r.GET("/employers", employerhandler.GetEmployers)
	r.GET("/jobseekers", jobseekerhandler.GetJobseekers)

	jobseeker := r.Group("/jobseeker")
	jobseeker.Use(middleware.JobSeekerMiddleware())
	{

		jobseeker.GET("/profile/:id", jobseekerhandler.GetJobseekerProfile)
		jobseeker.POST("/apply/job/:jobid", jobhandler.ApplyJob)
		jobseeker.GET("/applicants/:jobid", jobhandler.ListApplyJobsByid)

		jobseeker.POST("/follow/employer", jobseekerhandler.FollowEmployers)
		jobseeker.DELETE("/unfollow/employer", jobseekerhandler.UnFollowEmployers)
		jobseeker.GET("/follows", jobseekerhandler.GetFollowingEmployers)

		jobseeker.GET("/newsletter/:newsletterid", newletter.GetNewsLetter)
		jobseeker.GET("/newsletters/list", newletter.ListNewsLetters)
		jobseeker.POST("/newsletter/subscription", newletter.AddSubscription)
		jobseeker.PUT("/newletter/subscription/cancel/:subid", newletter.CancelSubscription)

		jobseeker.GET("/subscription", newletter.GetSubscription)
	}

	r.POST("/employer/signup", employerhandler.EmployerSignup)
	r.POST("/employer/login", employerhandler.LoginEmployer)

	employer := r.Group("/employer")
	employer.Use(middleware.EmployerMiddleware())
	{
		employer.POST("/job", jobhandler.CreateJob)
		employer.PUT("/job/:id", jobhandler.UpdateJob)

		employer.PUT("/applicant/accept/:jobid/:jobseekerid", jobhandler.AcceptApplicant)
		employer.PUT("/applicant/reject/:jobid/:jobseekerid", jobhandler.RejectApplicant)
		employer.GET("/applicant/:jobid/:status", jobhandler.GetApplicantsByStatus)

		employer.POST("/newsletter", newletter.CreateNewsletterService)
		employer.GET("/newsletter/subscribers/:newsletterid", newletter.GetSubscribers)
		employer.POST("/newsletter/send/:newsletterid", newletter.SendNewsletter)
	}

	return r
}

func Start(addr string) error {
	return r.Run(addr)
}
