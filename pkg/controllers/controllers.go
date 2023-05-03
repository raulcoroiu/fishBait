package controllers

import (
	fmt
)

func HashPassword (password string) string{
	fmt.Println(password)
}

func VerifyPassword(userPassword string, givenPassword) (bool, string){

}

func Singup() gin.HandleFunc{

	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeOut(context.Backgroud(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil{
			c.JSON{http.StatusBadRequest, gin.H{"error": err.Error()}}
			return
		}

		validationErr := Validate.Struct(user)
		if validationErr !=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
		}

		count, err := UserCollection.CountDocuments(ctx, bson.M{"error": user.Email})
		if err != nil{
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exist"})
		}

		count, err = User.Collection.CountDocuments(ctx, bson.M{"phone":user.Phone})

		defer cancel()
		if err != nil{
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error" : "this phone no is already used"})
		}
	}

	password := HashPassword(*user.Password)
	user.Password = &password

	user.Created_At, _ = time.Parse(time.RFC3339 , time.Now().Format(time.RFC3339))
	user.Updated_At, _ = time.Parse(time.RFC3339 , time.Now().Format(time.RFC3339))
	user.ID = primitve.NewObjectID()
	user.User_ID = user.ID.Hex() 
}

func Login() gin.HandleFunc{

}

func ProductViewAdmin() gin.HandleFunc{

}

func SearchProduct() gin.HandleFunc{

}

func SearchProductByQuery() gin.HandleFunc{

}