import React, { useState, useEffect } from "react";
import axios from "axios";
import Ranking from "./Ranking";

export interface Image {
  id: string;
  race: string;
  filename: string;
  good: number;
  nop: number;
  name: string;
}

const Top = () => {
  const [images, setImages] = useState<Image[]>([]);

  //非同期処理でDBから画像を取得
  useEffect(() => {
    axios.get("http://127.0.0.1:8080/reactGetImg").then((res) => {
      setImages(res.data);
    });
  }, []);
  return (
    <>
      <h1>Athlete App</h1>
      <p>美女ランキング</p>
      <div className="views">
        <div className="person">
          <img
            src="https://www.rikujyokyogi.co.jp/wp-content/uploads/2021/03/EN4A0178.jpg"
            alt=""
          />
          <Ranking images={images} />
        </div>
      </div>
      <p></p>
    </>
  );
};

export default Top;
