package auth

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	db      *mongo.Collection
	context context.Context
}

// The above type defines a repository interface for CRUD operations on a user entity with additional
// methods for reading by ID, email, and phone.
// @property Create - Create is a method that takes an input of type InUser and returns a User and an
// error. It is used to create a new user in the repository.
// @property Read - The Read method takes an ID as input and returns a User object and an error. It is
// used to retrieve a user from the repository based on their ID.
// @property Update - Update is a method defined in the Repository interface that takes in an id string
// and a map of fields to update as input parameters. It returns a User object and an error. The method
// is responsible for updating the user with the given id in the repository with the new values
// provided in the map. If
// @property {bool} Delete - Delete is a method of the Repository interface that takes an integer
// parameter representing the ID of a user and returns a boolean value indicating whether the user was
// successfully deleted from the repository or not.
// @property ReadByID - ReadByID is a method of the Repository interface that takes in a string
// parameter representing the ID of a user and returns a User object and an error. This method is used
// to retrieve a user from the repository by their ID.
// @property ReadByEmail - ReadByEmail is a method defined in the Repository interface that takes a
// string parameter "email" and returns a User and an error. This method is used to retrieve a user
// from the repository by their email address.
// @property ReadByPhone - ReadByPhone is a method defined in the Repository interface that takes a
// phone number as a string parameter and returns a User object and an error. This method is used to
// retrieve a user from the repository based on their phone number.
type Repository interface {
	Create(in InUser) (User, error)
	Read(id string) (User, error)
	Update(id string, upd map[string]interface{}) (User, error)
	Delete(id string) bool
	ReadByID(id string) (User, error)
	ReadAll() ([]User, error)
	ReadByEmail(email string) (User, error)
	ReadByPhone(phone string) (User, error)
}

// This is a method of the `Repo` struct that implements the `Create` function of the `Repository`
// interface. It takes an `InUser` object as input, converts it to a `User` object using the `ToUser`
// method of the `InUser` object, inserts the `User` object into the MongoDB collection using the
// `InsertOne` method of the `mongo.Collection` object, and returns the created `User` object along
// with any error that occurred during the insertion.
func (s *Repo) Create(in InUser) (User, error) {
	user := in.ToUser()
	_, err := s.db.InsertOne(s.context, user)
	if err != nil {
		return User{}, err
	}
	return *user, nil
}

// This is a method of the `Repo` struct that implements the `Read` function of the `Repository`
// interface. It takes an `id` string as input, retrieves the corresponding `User` object from the
// MongoDB collection using the `FindOne` method of the `mongo.Collection` object, decodes the
// retrieved document into a `User` object using the `Decode` method, and returns the `User` object
// along with any error that occurred during the retrieval.
func (s *Repo) Read(id string) (User, error) {
	var user User
	err := s.db.FindOne(s.context, id).Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// This is a method of the `Repo` struct that implements the `Update` function of the `Repository`
// interface. It takes an `id` string and a map of fields to update as input parameters, updates the
// corresponding `User` object in the MongoDB collection using the `FindOneAndUpdate` method of the
// `mongo.Collection` object, and returns the updated `User` object along with any error that occurred
// during the update. The `FindOneAndUpdate` method updates the document in the collection that matches
// the given `id` with the new values provided in the `upd` map. The updated document is then decoded
// into a `User` object using the `Decode` method and returned.
func (s *Repo) Update(id string, upd map[string]interface{}) (User, error) {
	var user User
	err := s.db.FindOneAndUpdate(s.context, id, upd).Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// This is a method of the `Repo` struct that implements the `Delete` function of the `Repository`
// interface. It takes an `id` string as input, deletes the corresponding `User` object from the
// MongoDB collection using the `DeleteOne` method of the `mongo.Collection` object, and returns a
// boolean value indicating whether the deletion was successful or not. If an error occurs during the
// deletion, the method returns `false`.
func (s *Repo) Delete(id string) bool {
	_, err := s.db.DeleteOne(s.context, id)
	if err != nil {
		return false
	}
	return true
}

// This is a method of the `Repo` struct that implements the `ReadByID` function of the `Repository`
// interface. It takes an `id` string as input, retrieves the corresponding `User` object from the
// MongoDB collection using the `FindOne` method of the `mongo.Collection` object with the `id` as the
// filter, decodes the retrieved document into a `User` object using the `Decode` method, and returns
// the `User` object along with any error that occurred during the retrieval. This method is used to
// retrieve a user from the repository by their ID.
func (s *Repo) ReadByID(id string) (User, error) {
	var user User
	err := s.db.FindOne(s.context, id).Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// This is a method of the `Repo` struct that implements the `ReadByEmail` function of the `Repository`
// interface. It takes an `email` string as input, retrieves the corresponding `User` object from the
// MongoDB collection using the `FindOne` method of the `mongo.Collection` object with the `email` as
// the filter, decodes the retrieved document into a `User` object using the `Decode` method, and
// returns the `User` object along with any error that occurred during the retrieval. This method is
// used to retrieve a user from the repository by their email address.
func (s *Repo) ReadByEmail(email string) (User, error) {
	var user User
	err := s.db.FindOne(s.context, email).Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// This is a method of the `Repo` struct that implements the `ReadByPhone` function of the `Repository`
// interface. It takes a `phone` string as input, retrieves the corresponding `User` object from the
// MongoDB collection using the `FindOne` method of the `mongo.Collection` object with the `phone` as
// the filter, decodes the retrieved document into a `User` object using the `Decode` method, and
// returns the `User` object along with any error that occurred during the retrieval. This method is
// used to retrieve a user from the repository based on their phone number.
func (s *Repo) ReadByPhone(phone string) (User, error) {
	var user User
	err := s.db.FindOne(s.context, phone).Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *Repo) ReadAll() ([]User, error) {
	var users []User
	cursor, err := s.db.Find(s.context, bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(s.context, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (s *Repo) ReadByPhoneOrEmail(phone string, email string) (User, error) {
	var user User
	err := s.db.FindOne(s.context, bson.M{"$or": []bson.M{{"phone": phone}, {"email": email}}}).Decode(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// The function returns a new instance of a Repository interface implementation with a MongoDB database
// connection.
func NewRepo(db *mongo.Database) Repository {
	ctx := context.TODO()
	return &Repo{db: db.Collection("users"), context: ctx}
}
