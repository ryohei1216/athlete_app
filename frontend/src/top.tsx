import axios from "axios";
import React, { useEffect, useState } from "react";
import { Image } from "./App";

interface Props {
  images: Image[];
}

const Top: React.VFC<Props> = (props) => {
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
          <p>テキストテキストテキストテキストテキストテキスト</p>
          <p>テキストテキストテキストテキストテキストテキスト</p>
        </div>
      </div>
      <p>{props.images}</p>
    </>
  );
};

export default Top;
