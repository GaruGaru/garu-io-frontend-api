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
		Role:      "Lead Platform developer",
		StartYear: 2018,
		EndYear:   -1,
		Company: models.Company{
			Name:    "Data Appeal",
			Website: "https://datappeal.io/",
			Logo:    "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAOEAAADhCAMAAAAJbSJIAAAAYFBMVEUiJjL///8fIzBAQ00bIC1hZGuEhYoVGikABhzFxsjx8fIyNUAAABMADB/u7u63uLsAABbj4+Spq69SVV2foKMMEyOQkpY5PEdcXmaXmZ2vsbSlpqgAAAD4+PgAABvW1tjc68prAAABv0lEQVR4nO3cwVLbMBSGUTelBGhpgMbQQinv/5bstUBzx47wn55vL+mereTx9Hh53h2np/1ZN/+evu6mc+7iG2F6hPkR5keYH2F+hPkR5keYH2F+hPkR5keYH2F+hPkR5keYH2F+hPkR5keYH2F+hPkR5keYH2F+hPkR5keYH2F+hPmtIJx7Ffdvl++XjbdcOP95+Ljn2gHzsVl/uYy4XHh3/6XTr4vKfj9umuXXV4vm26DwJ2Etwn6EhG2E1Qj7ERK2EVYj7EdI2EZYjbAfIWEbYTXCfoSEbYTVCPsRErZtT7h/PH7cy1PpgO0Jp/26r9wbFK4cYTXC8RFWIxwfYTXC8RFWIxwfYTXC8RFWIxwfYTXC8RFWIxwfYTXC8RFWIxwfYTXC8RFWIxwfYTXC8RFWIxwfYbU1/k9zu2p/NyecXw6r9vy6NWH/C9qFERISEhISEhISEhISnoXw39vNSXs93H6ycPf91C2bb4VbjN2pWzbe9m6i1o4wP8L8CPMjzI8wP8L8CPMjzI8wP8L8CPMjzI8wP8L8CPMjzI8wP8L8CPMjzI8wP8L8CPMjzI8wP8L8CPMjzI8wP8L8/gfhOx1iXMi+eXo4AAAAAElFTkSuQmCC",
		},
		Description: `
<p><p>During my experience as fullstack developer in TravelAppeal I've worked on the design and implementation of the <b>microservice architecture</b>
using technologies such as <b>spring boot</b> for the backend and <b>angular/typescript</b> for the frontend deployed on <b>AWS (ecs)</b> and
<b>Azure (kubernetes)</b>.<p><p>Design and implementation of <b>high scalable data ingestion pipeline</b> and <b>ETL jobs</b> on AWS using <b>Amazon RDS, 
Amazon Redshift and Amazon S3</b> as backend storages and <b>ECS</b> as container orchestration platform with Amazon 
<b>CodeBuild/CodePipeline for CI/CD</b> 

Migration from ECS to EKS (Managed kubernetes service) using Terraform for infrastructure provisioning, providing reliable multi cluster, docker based CI/CD pipeline using CodeBuild design and deployment of monitoring and real time alerting infrastructure for aws/eks using Prometheus .
<p></p>

Implementation of the internal developer platform using Backstage to improve dev productivity, enable automatization and archive faster and safer development iterations. 

Public contributions to widely used open source projects such as TrinoDB and Jaeger. 
`,
	},
	{
		Role:      "Full Stack Developer",
		StartYear: 2016,
		EndYear:   2018,
		Company: models.Company{
			Name:    "Instal",
			Website: "https://instal.com/",
			Logo:    "https://media-exp1.licdn.com/dms/image/C560BAQHeqaAqYFphcg/company-logo_200_200/0?e=2159024400&v=beta&t=fdWnWI221qXh3u3dApKhYRPzmSk_hdnqp70jwZCoD5Y",
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
