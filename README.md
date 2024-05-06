# template-serverless-service

A Lambda and other infrastructure to track direct-from-S3 downloads of Discover datasets. That is, code responsible for
populating the rows in the `dataset_downloads` with `origin == AWSRequesterPayer` table in the Discover Postgres
database.

