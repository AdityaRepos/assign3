import React, { useState } from 'react';
import Axios from 'axios';
import './App.css';

const baseURL = "http://localhost:8080/posts";

function App() {
  const [post, setPost] = useState({ ID: "", Post: "" });
  const [allPosts, setAllPosts] = useState([]);

  function handlePost() {
    Axios.post(baseURL, post)
      .then((response) => {
        console.log(response.data);
        setPost({ ID: "", Post: "" });
      })
      .catch((error) => {
        console.error(error);
      });
  }

  function handleDelete() {
    Axios.delete(`${baseURL}/${post.ID}`)
      .then((response) => {
        console.log(response.data);
        setPost({ ID: "", Post: "" });
      })
      .catch((error) => {
        console.error(error);
      });
  }

  function handleGet() {
    Axios.get(`${baseURL}/${post.ID}`)
      .then((response) => {
        setPost(response.data);
        setAllPosts([]);
      })
      .catch((error) => {
        console.error(error);
      });
  }

  function handleGetAll() {
    Axios.get(baseURL)
      .then((response) => {
        setAllPosts(response.data);
        setPost({ ID: "", Post: "" });
      })
      .catch((error) => {
        console.error(error);
      });
  }

  return (
    <div className="App">
      <header className="App-header">
        <textarea
          value={post.ID}
          id="ID"
          rows="1"
          cols="30"
          onChange={(e) => setPost({ ...post, ID: e.target.value })}
        />
        <textarea
          value={post.Post}
          id="Post"
          rows="3"
          cols="40"
          onChange={(e) => setPost({ ...post, Post: e.target.value })}
        />
        <button id="Post" onClick={handlePost}>
          Post
        </button>
        <button id="Delete" onClick={handleDelete}>
          Delete
        </button>
        <button id="Get" onClick={handleGet}>
          Get
        </button>
        <button id="GetAll" onClick={handleGetAll}>
          All Posts
        </button>
      </header>
      <p>
        {post.ID !== "" ? (
          <strong>
            The Post -&gt; ID: {post.ID}, Post: {post.Post}
          </strong>
        ) : (
          <div>
            <strong>All Posts:</strong>
            <table>
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Post</th>
                </tr>
              </thead>
              <tbody>
                {allPosts.map((post) => (
                  <tr key={post.ID}>
                    <td>{post.ID}</td>
                    <td>{post.Post}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </p>
    </div>
  );
}

export default App;
