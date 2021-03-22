# CMDB Docs

## CMDB Data Model

The process of designing the data model for the CMDB application could be
iterative and maybe use visual tools. I created a sample ./model directory
here with a make target for generating ERD:
```sh
make erd
Create file://doc/db/out.html and file://doc/db/out.pdf
```
Which creates targets e.g. [cmdb ERD](db/out.pdf)
that you can view.

We can now start to design the proto definiton of our data model
using the above model defintion for the messsage and service
definitions.
[Refer to the proto buff defintion](https://github.com/infobloxopen/protoc-gen-gorm)
for database defintion which is layered on top of [Gorm](http://gorm.io/).

We would build out the data model incrementaly starting with the top node
in our case the Region. Then add the necessary migration for it, see below
for more detail on database migration below.

Once the database was completed I created an ERD using LucidChart import feature
and generated a more complete ERD, this process is tedious so not something for fast
iterative process. There are two ERDs one for the data model to drive
[Application configuration based on Helm Charts](db/cmdb_app_config_erd.pdf).
The other is geared toward the
[Application Deployment model based on Kubernetes](db/cmdb_app_deployment_erd.pdf).

Note you could start with ERD diagram, but there is not an easy way
to go from ERD to create the protobuf definition right now.
