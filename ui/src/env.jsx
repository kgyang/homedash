import { useState } from 'react'
import axios from 'axios'
import './env.css'

function Env() {
  const [reading, setReading] = useState(false)
  const [envdata, setEnvData] = useState(null)

  const readEnvData = async () => {
    try {
      setReading(true)
      const response = await axios.get(window.location.toString() + 'api/v1/env');
      const ret = response.data;
      console.log("env", ret);
      setEnvData(ret);
      setReading(false);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <>
      <div id="env">
        <button onClick={readEnvData} disabled={reading}>
	  {envdata ? "Click to refresh environment data":"Click to get environment data"}
        </button>
        {envdata && envdata.temperature && <p className="temperature">Temperature<br/>{envdata.temperature}</p> }
        {envdata && envdata.humidity && <p className="humidity">Humidity<br/>{envdata.humidity}</p> }
        {envdata && envdata.ch2o && <p className="ch2o">CH2O<br/>{envdata.ch2o}</p> }
      </div>
    </>
  )
}

export default Env
