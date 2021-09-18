import React, { useEffect, useState } from "react";
import Top from "./top";
import axios from "axios";

export interface Image {
  id: string;
  race: string;
  filename: string;
  good: number;
  nop: number;
  name: string;
}

function App() {
  let images: Image[] = [];
  //非同期処理でDBから画像を取得
  useEffect(() => {
    axios.get("http://127.0.0.1:8080/reactGetImg").then((res) => {
      images = res.data;
      console.log(images);
    });
  }, []);
  return (
    <div className="App">
      <header className="App-header">
        <p>This is App.tsx</p>
      </header>
      <Top images={images} />
    </div>
  );
}

export default App;
