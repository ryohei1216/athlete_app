import React, { useState } from "react";
import { storage } from "./firebase/firebase";
import { Image } from "./Top";
import axios from "axios";

const Upload = () => {
  const image: Image = {
    id: "",
    race: "",
    filename: "",
    good: 0,
    nope: 0,
    name: "",
    url: "",
  };
  const [file, setFile] = useState<File>();
  const [name, setName] = useState<string>("");
  const [race, setRace] = useState<string>("");

  const uploadFile: (param: any) => void = (event: any) => {
    event.preventDefault();

    if (file) {
      image.name = name;
      image.race = race;
      image.filename = file.name;

      // Create a root reference
      var storageRef = storage.ref();
      // Create a reference to 'mountains.jpg'
      var mountainsRef = storageRef.child("images/" + file.name);
      mountainsRef.put(file).then(function (snapshot) {
        console.log("Uploaded a blob or file!");
        snapshot.ref.getDownloadURL().then((downloadURL) => {
          console.log(downloadURL);
          image.url = downloadURL;
        });
        //GoのDBに保存
        console.log(image);
        axios.defaults.headers.post["Content-Type"] =
          "application/json;charset=utf-8";
        axios.defaults.headers.post["Access-Control-Allow-Origin"] = "*";
        axios
          .post("http://127.0.0.1:8080/reactUploading", { image })
          .then((res) => {});
      });
    }
  };

  return (
    <div>
      <form onSubmit={uploadFile}>
        <input
          type="file"
          onChange={(e) => {
            if (e.target.files !== null) {
              setFile(e.target.files[0]);
            }
          }}
        />
        <input
          type="text"
          name="name"
          onChange={(e) => {
            setName(e.target.value);
          }}
        />
        選手名
        <input
          type="radio"
          name="race"
          value="asian"
          onChange={(e) => {
            setRace(e.target.value);
          }}
        />
        アジア系
        <input
          type="radio"
          name="race"
          value="white"
          onChange={(e) => {
            setRace(e.target.value);
          }}
        />
        白人
        <input
          type="radio"
          name="race"
          value="black"
          onChange={(e) => {
            setRace(e.target.value);
          }}
        />
        黒人
        <input type="submit" />
      </form>
    </div>
  );
};

export default Upload;
