package auth

// The `import` statement is used to import packages that are required for the code to run. In this
// case, the code is importing the `http` package, the `auth` package from the `sharir/pkg` directory,
// and the `fiber` package from the `github.com/gofiber/fiber/v2` repository. These packages are used
// in the code to handle HTTP requests and responses, and to implement user authentication
// functionality.
import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthBody struct {
	PhoneNumber string `json:"phonenumber"`
	Password    string `json:"password"`
}

// The above type is a struct representing a user with various fields such as ID, name, password, phone
// number, profile picture, email, username, user type, and creation time.
// @property {string} ID - A unique identifier for the user, typically stored as a string.
// @property {string} Name - Name is a property of the User struct that represents the name of the
// user.
// @property {string} Password - The password property is a string that stores the user's password. It
// is important to ensure that this property is properly secured and encrypted to prevent unauthorized
// access to the user's account.
// @property {string} PhoneNumber - The phone number of the user.
// @property {string} ProfilePic - The ProfilePic property is a string that represents the URL or file
// path of the user's profile picture.
// @property {string} Email - The email property is a string that represents the email address of a
// user. It is used to uniquely identify a user and to communicate with them via email.
// @property {string} Username - The username property is a string that represents the unique username
// of a user. It is used for authentication and identification purposes.
// @property {string} UserType - UserType is a property of the User struct that represents the type of
// user. It can be used to differentiate between different types of users, such as regular users,
// admins, or moderators. The value of UserType can be set to any string value that represents the type
// of user.
// @property CreatedAt - CreatedAt is a property of the User struct that represents the date and time
// when the user was created. It is of type time.Time and is formatted as "YYYY-MM-DD HH:MM:SS". This
// property can be used to track when a user was added to the system and to sort users by
type User struct {
	ID          string    `json:"id" bson:"_id"`
	Name        string    `json:"name"`
	Password    string    `json:"password"`
	PhoneNumber string    `json:"phone_number"`
	ProfilePic  string    `json:"profile_pic"`
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	UserType    string    `json:"usertype"`
	CreatedAt   time.Time `json:"created_at"`
}

// The above type defines the structure of an input user object in Go, including fields for name,
// password, phone number, profile picture, email, username, and user type.
// @property {string} Name - The name of the user.
// @property {string} Password - The "Password" property is a string that represents the user's
// password. It is likely used for authentication purposes to ensure that only authorized users can
// access certain features or information. It is important to ensure that passwords are securely stored
// and encrypted to prevent unauthorized access.
// @property {string} PhoneNumber - The phone number property is a string that represents the user's
// phone number. It is included in the InUser struct as a JSON field with the key "phonenumber".
// @property {string} ProfilePic - ProfilePic is a property of the InUser struct that represents the
// URL or file path of the user's profile picture. It is of type string and is tagged with
// `json:"profilepic"` to indicate that it should be marshaled and unmarshaled as a JSON field with the
// key "profile
// @property {string} Email - The email property is a string that represents the email address of a
// user. It is used to uniquely identify a user and to send notifications or messages to them.
// @property {string} Username - The username property is a string that represents the unique
// identifier for a user's account. It is often used as a way for users to log in to a system or
// application.
// @property {string} UserType - UserType is a property of the InUser struct that represents the type
// of user. It can be used to differentiate between different types of users, such as regular users,
// administrators, or moderators. The value of UserType can be set to any string value that represents
// the type of user.
type InUser struct {
	Name        string `json:"name"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phonenumber"`
	ProfilePic  string `json:"profilepic"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	UserType    string `json:"usertype"`
}

// The above type defines the structure of an output user object with various properties such as ID,
// name, email, phone number, profile picture, user type, username, and creation timestamp.
// @property {string} ID - The unique identifier for the user, typically stored as a string.
// @property {string} Name - Name is a string property that represents the name of a user.
// @property {string} Email - The email address of the user.
// @property {string} PhoneNumber - The phone number of the user.
// @property {string} ProfilePic - ProfilePic is a property of the OutUser struct that represents the
// URL or file path of the user's profile picture.
// @property {string} UserType - UserType is a property of the OutUser struct that represents the type
// of user. It could be a customer, admin, or any other type of user.
// @property {string} Username - The username property is a string that represents the unique username
// of a user. It is used for authentication and identification purposes.
// @property CreatedAt - CreatedAt is a property of the OutUser struct that represents the date and
// time when the user was created. It is of type time.Time and is formatted as "YYYY-MM-DD HH:MM:SS".
type OutUser struct {
	ID          string    `json:"id" bson:"_id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	ProfilePic  string    `json:"profile_pic"`
	UserType    string    `json:"user_type"`
	Username    string    `json:"username"`
	CreatedAt   time.Time `json:"created_at"`
}

// `func (in *InUser) ToUser() User` is a method defined on the `InUser` struct that converts an
// `InUser` object to a `User` object. It creates a new `User` object with the same values as the
// `InUser` object, except for the `ID` and `CreatedAt` fields. The `ID` field is generated using the
// `uuid` package, and the `CreatedAt` field is set to the current time using the `time` package. The
// `Password` field is hashed using the `hashPassword` function defined in the same file. This method
// is useful when we want to create a new user object from user input data, such as when a user
// registers for a new account.
func (in *InUser) ToUser() User {
	uuid := uuid.New().String()
	return User{
		ID:          uuid,
		Name:        in.Name,
		ProfilePic:  in.ProfilePic,
		PhoneNumber: in.PhoneNumber,
		Password:    hashPassword(in.Password),
		Email:       in.Email,
		UserType:    in.UserType,
		Username:    in.Username,
		CreatedAt:   time.Now(),
	}
}

// `func (u *User) ToOutUser() OutUser` is a method defined on the `User` struct that converts a `User`
// object to an `OutUser` object. It creates a new `OutUser` object with the same values as the `User`
// object, and returns it. This method is useful when we want to return a user object to the client in
// a response, but we don't want to expose all the fields of the `User` object. Instead, we can create
// a new `OutUser` object with only the fields we want to expose, and return that to the client.
func (u *User) ToOutUser() OutUser {
	return OutUser{
		ID:          u.ID,
		Name:        u.Name,
		ProfilePic:  u.ProfilePic,
		PhoneNumber: u.PhoneNumber,
		UserType:    u.UserType,
		Email:       u.Email,
		Username:    u.Username,
		CreatedAt:   u.CreatedAt,
	}
}

// The function takes a password string, generates a hash using bcrypt algorithm with minimum cost, and
// returns the hash as a string.
func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes)
}
