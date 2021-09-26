import firebase from "firebase/compat/app";
import "firebase/compat/storage";

var firebaseConfig = {
  apiKey: "AIzaSyCVteTDGH9rkIyTRy39pMk99lNhgGHn66s",
  authDomain: "athlete-app-front.firebaseapp.com",
  projectId: "athlete-app-front",
  storageBucket: "athlete-app-front.appspot.com",
  messagingSenderId: "931400101738",
  appId: "1:931400101738:web:652ffbe4531d3dd382ff13",
  measurementId: "G-9Z7CY96FDF",
};

export var firebaseApp = firebase.initializeApp(firebaseConfig);
export var storage = firebase.storage();
