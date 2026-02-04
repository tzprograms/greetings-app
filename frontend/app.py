import streamlit as st
import requests
import os

st.title("Kubernetes Microservices Demo")


backend_url = os.getenv("BACKEND_URL", "http://go-backend-service:8080/api/data")

st.write("Click the button to fetch data from the Go API:")

if st.button('Fetch Data'):
    try:
        response = requests.get(backend_url)
        data = response.json()
        
        # 3. Display the result
        st.success(f"Backend says: {data['text']}")
    except Exception as e:
        st.error(f"Could not connect to backend: {e}")