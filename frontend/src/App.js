import {useState} from 'react';
import axios from 'axios';

function App()
{
  const [url, setUrl] = useState('');
  const [response, setResponse] = useState('');
  // `http://localhost:8080/api/parse?q=${url}`

  const fetchData = () => 
  {
    console.time();
    fetch(`http://localhost:8080/api/parse?q=${url}`)
      .then(res => 
      {
        res.text()
          .then(res => 
          {
            // setResponse(res)
            console.timeEnd();
            console.log(res);
          });
      });
  };

  return (
    <div
      className='App'
      style={{
        display: 'flex',
        flexDirection: 'column',
        maxWidth: '300px',
        maxHeight: '100px',
      }}
    >
      <textarea
        name='url'
        id=''
        cols='30'
        rows='10'
        style={{
          marginBottom: '1rem',
        }}
        onChange={e =>
        {
          e.preventDefault();
          setUrl(e.currentTarget.value.replaceAll("/", "%2F").replaceAll(":", "%3A"));
        }}
      />
      <button
        onClick={e =>
        {
          e.preventDefault();
          fetchData();
        }}
      >
        Submit
      </button>
      <p>{response}</p>
    </div>
  );
}

export default App;
