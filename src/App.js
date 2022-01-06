import React, { useState, useEffect } from "react";

const server = "http://localhost:8081/users";

const App = () => {
  const [posts, setPosts] = useState([]);

  useEffect(() => {
    fetch(server, {
      method: "GET",
      mode: "cors",
    })
      .then((res) => res.json())
      .then((data) => {
        setPosts(data);
      });
  }, []);
  return (
    <div>
      <ul>
        {posts.map((post) => (
          <div>
            <li key={post.id}> Name:{post.name}</li>
            <li key={post.id}> Price:{post.price}</li>
          </div>
        ))}
      </ul>
    </div>
  );
};

export default App;
