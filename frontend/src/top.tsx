import React, { useState, useEffect } from "react";
import axios from "axios";
import Ranking from "./Ranking";
import Upload from "./Upload";
import { storage } from "./firebase/firebase";

export interface Image {
  id: string;
  race: string;
  filename: string;
  good: number;
  nope: number;
  name: string;
  url: string;
}

const Top = () => {
  const [images, setImages] = useState<Image[]>([]);
  //imagesにimageObjのセットをする
  const [imageURL, setImageURL] = useState<string>("");

  //非同期処理でDBから画像を取得
  useEffect(() => {
    // axios.get("http://127.0.0.1:8080/reactGetImg").then((res) => {
    //   setImages(res.data);
    //   Dbからデータを受け取って中のURLデータを利用
    //
    // });
    var imageObj = storage.ref().child("download2.jpg");
    imageObj
      .getDownloadURL()
      .then((url) => {
        setImageURL(url);
      })
      .catch((error) => {
        console.log(error.message);
      });
  }, []);
  return (
    <>
      <h1>Athlete App</h1>
      <p>美女ランキング</p>
      <Upload />
      <div className="views">
        <div className="person">
          <img src={imageURL} alt="" />
          <Ranking images={images} />
        </div>
      </div>
    </>
  );
};

export default Top;
