package main

import "github.com/bartoszpiechnik25/static-files-server/server"

func main() {
	// sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	// if err != nil {
	// 	fmt.Println("Could not load default configuration!")
	// 	return
	// }
	// fmt.Println(sdkConfig.Region)

	// s3Client := s3.NewFromConfig(sdkConfig)

	// buckets, err := s3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	// if err != nil {
	// 	fmt.Println("Could not list buckets for this account!")
	// 	return
	// }

	// if len(buckets.Buckets) == 0 {
	// 	fmt.Println("You don't have any buckets!")
	// } else {
	// 	for _, bucket := range buckets.Buckets {
	// 		fmt.Printf("Your bucker name: %s\n", *bucket.Name)
	// 	}
	// }

	// router := gin.Default()

	// router.GET("/hello", func(ctx *gin.Context) {
	// 	if bearer, ok := ctx.Request.Header["Authorization"]; ok {
	// 		fmt.Println(bearer)
	// 	} else {
	// 		fmt.Println("No header!")
	// 		ctx.JSON(http.StatusUnauthorized, "unathorized")
	// 		ctx.Abort()
	// 	}
	// 	ctx.Next()
	// })

	// router.Run(":42069")
	router := server.CreateServer()
	router.Run(":6666")
}
