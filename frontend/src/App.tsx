import React, { useEffect, useState } from "react";
import Top from "./top";
import axios from "axios";

interface Image {
  id: string;
  race: string;
  filename: string;
  good: number;
  nop: number;
  name: string;
}

function App() {
  const [images, setImages] = useState<Image[]>([]);

  useEffect(() => {
    axios.get("http://127.0.0.1:8080/reactGetImg").then((res) => {
      setImages(res.data[0]);
      console.log(res.data);
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
