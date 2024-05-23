package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	"github.com/tmpmadula/cantina-shop/internal/models"

	"github.com/go-playground/validator/v10"
)

// Initialize the validator
var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Custom validation errors
var (
	ErrInvalidEmail    = errors.New("invalid email format")
	ErrInvalidPassword = errors.New("password must be at least 8 characters long and contain a number")
)

// ValidateUser validates user registration data
func ValidateUser(user *models.User) error {
	if err := validate.Struct(user); err != nil {
		return err
	}
	if !isValidEmail(user.Email) {
		return ErrInvalidEmail
	}
	if !isValidPassword(user.Password) {
		return ErrInvalidPassword
	}
	return nil
}

// ValidateDish validates dish data
func ValidateDish(dish *models.Dish) error {
	return validate.Struct(dish)
}

// ValidateDrink validates drink data
func ValidateDrink(drink *models.Drink) error {
	return validate.Struct(drink)
}

// ValidateReview validates review data
func ValidateReview(review *models.Review) error {
	if review.Rating < 1 || review.Rating > 5 {
		return errors.New("rating must be between 1 and 5")
	}
	return validate.Struct(review)
}

// Middleware to validate user data
func UserValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := ValidateUser(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Middleware to validate dish data
func DishValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var dish models.Dish
		if err := json.NewDecoder(r.Body).Decode(&dish); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := ValidateDish(&dish); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Middleware to validate drink data
func DrinkValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var drink models.Drink
		if err := json.NewDecoder(r.Body).Decode(&drink); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := ValidateDrink(&drink); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Middleware to validate review data
func ReviewValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var review models.Review
		if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := ValidateReview(&review); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Helper functions
func isValidEmail(email string) bool {
	// Regular expression to validate email format
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func isValidPassword(password string) bool {
	// Check if password is at least 8 characters long and contains at least one number
	if len(password) < 8 {
		return false
	}
	re := regexp.MustCompile(`[0-9]`)
	return re.MatchString(password)
}
