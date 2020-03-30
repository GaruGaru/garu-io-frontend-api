package repository

import "garu-io-frontend-api/models"

// If you want to hire me and you are reading this: forgive me, but I'm not going to setup any kind of storage in order
// to write 5 records on it.

var knownLanguages = []models.Languages{
	{
		Name:       "Java",
		Experience: "Experienced",
		Level:      98,
	},
	{
		Name:       "Kotlin",
		Experience: "expert",
		Level:      83,
	},
	{
		Name:       "GoLang",
		Experience: "Proficient",
		Level:      72,
	},
	{
		Name:       "Python",
		Experience: "Proficient",
		Level:      61,
	},
	{
		Name:       "Typescript",
		Experience: "Beginner",
		Level:      40,
	},
}

var workExperiences = []models.WorkExperience{
	{
		Role:      "Full Stack Developer",
		StartYear: 2018,
		EndYear:   -1,
		Company: models.Company{
			Name:    "Data Appeal",
			Website: "https://datappeal.io/",
			Logo:    "https://media-exp1.licdn.com/dms/image/C4D0BAQEcpvYxu3dLyw/company-logo_100_100/0?e=1593648000&v=beta&t=9-ak9-W-ix97VeNAGu7BT_eXEiUCA22WD6TkDFwWImw",
		},
		Description: `
<p><p>During my experience as fullstack developer in TravelAppeal I've worked on the design and implementation of the <b>microservice architecture</b>
using technologies such as <b>spring boot</b> for the backend and <b>angular/typescript</b> for the frontend deployed on <b>AWS (ecs)</b> and
<b>Azure (kubernetes)</b>.<p><p>Design and implementation of <b>high scalable data ingestion pipeline</b> and <b>ETL jobs</b> on AWS using <b>Amazon RDS, 
Amazon Redshift and Amazon S3</b> as backend storages and <b>ECS</b> as container orchestration platform with Amazon 
<b>CodeBuild/CodePipeline for CI/CD</b> 

Migration from ECS to EKS (Managed kubernetes service) using Terraform for infrastructure provisioning, providing reliable multi cluster, docker based CI/CD pipeline using CodeBuild design and deployment of monitoring and real time alerting infrastructure for aws/eks using Prometheus .
<p></p>`,
	},
	{
		Role:      "Full Stack Developer",
		StartYear: 2018,
		EndYear:   -1,
		Company: models.Company{
			Name:    "Instal",
			Website: "https://instal.com/",
			Logo:    "https://media-exp1.licdn.com/dms/image/C560BAQHeqaAqYFphcg/company-logo_100_100/0?e=1593648000&v=beta&t=g4X_-y_ZCd8wdM1J4PWUxwRczoJO1ifEsmQMTW1fFMo",
		},
		Description: `
<p><p>In Instal I've worked as an android developer, my work was focused mainly on the proprietary <b>sdk development</b></p> <p> 
As Backend Developer I've worked with <b>apache storm</b> and <b>google dataflow</b> with <b>hadoop</b> in order to archive high performance real data processing, 
I've also contributed to the microservices-based architecture by developing scalable and fault tolerant services using <b>python</b> with <b>tornado</b> and 
<b>golang</b> with <b>gingonic</b>. </p>  <p> As frontend developer I've developed dashboards for our services using <b>angular 6</b> and <b>reactjs</b>.
</p> <p>During my experience in Instal I've gained experience also with the devops stack working with <b>Chef</b> (infrastructure as code), 
<b>Docker</b> (container engine), <b>Kubernetes</b> for production container orchestration and <b>Sensu</b> for monitoring and alerting system 
in the <b>GCE</b> environment.</p></p>
`,
	},
}
