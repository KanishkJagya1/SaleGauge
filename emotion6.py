from vaderSentiment.vaderSentiment import SentimentIntensityAnalyzer
import speech_recognition as sr
# import nltk
from textblob import TextBlob as blob

recognizer = sr.Recognizer()

with sr.Microphone() as source:
    print("Clearing background noise...")
    print('Say Something...')
    recorded_audio = recognizer.listen(source, timeout=1)
    print("Listening...")
    try:
        text = recognizer.recognize_google(recorded_audio)
        tb = blob(text)
        # print(text)
        print(tb.tags)
        print(tb.sentiment)
    except Exception as ex:
        print('Sorry... Try again')