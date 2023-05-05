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
	token, refreshtoken, _ := generate.TokenGenerator(*user.Email, *user.Firs_Name, *user.Last_Name, *user.User_ID)
	user.Token = &token 
	user.Refresh_Token = &refreshtoken
	user.UserCart = make([]moleds.ProductUser, 0)
	user.Adress_Details = make([]models.Adress, 0)
	user.Order_Status = make([]models.Order, 0)
	UserCollection.InsertOne(ctx, user)
	_, inserterr := UserCollection.InsertOne(ctx, user)
	if inserterr != nil{
		x.JSON(http.StatusInternalServerError, gin.H{"error":"the user did not get created"})
		return
	}
	defer cancel()

	c.JSON(http.StatusCreated, "Successfully signed in!")
}

func Login() gin.HandleFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeOut(context.Backgroud(). 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		err := UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&founduser)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "login or password incorrect"})
		}
	}

}

func ProductViewAdmin() gin.HandleFunc{

}

func SearchProduct() gin.HandleFunc{

}

func SearchProductByQuery() gin.HandleFunc{

}