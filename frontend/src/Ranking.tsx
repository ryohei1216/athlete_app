import React from "react";
import { Image } from "./Top";

interface Props {
  images: Image[];
}

const Ranking = (props: Props) => {
  let imageList: JSX.Element[] = [];
  if (props.images != null) {
    imageList = props.images.map((image) => (
      <li key={image.id}>
        <div>
          <p>{image.name}</p>
          <p>{image.good}</p>
          <p>{image.nope}</p>
        </div>
      </li>
    ));
  }
  return <div>{<ul>{imageList}</ul>}</div>;
};

export default Ranking;
