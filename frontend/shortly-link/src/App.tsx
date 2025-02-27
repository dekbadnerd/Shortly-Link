import React, { useState, useEffect } from "react";
import axios from 'axios';

const App: React.FC = () => {
  const [url, setUrl] = useState('');
  const [shortUrl, setShortUrl] = useState('');
  const BASE_URL = "http://localhost:8000";

  useEffect(() => {
    document.title = "Shorten URL";
  }, []);

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    try {
      const response = await axios.post<{ short_url: string }>(`${BASE_URL}/shorten`, { url });
      const { short_url } = response.data;
      setShortUrl(short_url);
    } catch (error) {
      console.error(error);
    }
  };

  const handleOpenUrl = () => {
    if (shortUrl) {
      window.open(`http://127.0.0.1:8000/${shortUrl}`, '_blank');
    }
  };

  const handleCopyUrl = () => {
    if (shortUrl) {
      navigator.clipboard.writeText(`http://127.0.0.1:8000/${shortUrl}`);
    }
  };

  return (
    <div>
      <h1>Shorten URL</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="url"
          value={url}
          onChange={(event) => setUrl(event.target.value)}
          placeholder="Enter URL"
          required
        />
        <button type="submit">SHORTEN !!!</button>
      </form>
      {shortUrl && (
        <div>
          <div>
            <p>Shortened URL: <a href={`http://127.0.0.1:8000/${shortUrl}`} target="_blank" rel="noopener noreferrer">{`http://127.0.0.1:8000/${shortUrl}`}</a></p>
          </div>
          <button className="smallBtn" onClick={handleOpenUrl}>Open URL</button>
          <button className="smallBtn" onClick={handleCopyUrl}>Copy URL</button>
        </div>
      )}
    </div>
  );
};

export default App;