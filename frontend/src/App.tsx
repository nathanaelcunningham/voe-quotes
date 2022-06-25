import { useState } from "react";
import "./App.css";

function App() {
  const [data, setData] = useState<string | undefined>();

  return (
    <div className="App">
      <pre>{data}</pre>
      <button
        onClick={() =>
          fetch("http://localhost:3000").then(async (res) =>
            setData(await res.text())
          )
        }
      >
        Get Data
      </button>
    </div>
  );
}

export default App;
