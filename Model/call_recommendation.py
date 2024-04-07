from flask import Flask, request, jsonify
import pickle
import os

# Load the recommendation function from the .sav file
def load_recommendation_function(sav_file_path):
    with open(sav_file_path, 'rb') as f:
        return pickle.load(f)

# Initialize Flask application
app = Flask(__name__)

# Define route for credit card recommendations
@app.route('/recommend', methods=['POST'])
def recommend_credit_cards():
    # Get input text from the request
    input_text = request.json.get('input_text', '')
    
    # Call the recommendation function to get recommendations
    recommendations = recommend_credit_cards_function(input_text)
    
    # Return the recommendations as JSON response
    return jsonify({'recommendations': recommendations})

if __name__ == '__main__':
    # Construct the file path for the .sav file
    sav_file_path = os.path.join("Model", "recommend_credit_cards.sav")
    
    # Load the recommendation function
    recommend_credit_cards_function = load_recommendation_function(sav_file_path)
    
    # Run the Flask app
    app.run(debug=True)
