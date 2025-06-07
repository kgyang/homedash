import { useState } from 'react'
import axios from 'axios'
import './env.css'

function Env() {
  const [envdata, setEnvData] = useState(null)

  const readEnvData = async () => {
    try {
      const response = await axios.get('http://localhost:8080/api/v1/env');
      const ret = response.data;
      console.log("env", ret);
      setEnvData(ret);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <>
      <div className="env">
        <button onClick={readEnvData}>
          Click to fresh
        </button>
        {envdata && envdata.temperature && <p className="temperature">T: {envdata.temperature} </p> }
        {envdata && envdata.humidity && <p className="humidity">H: {envdata.humidity} </p> }
        {envdata && envdata.ch2o && <p className="ch2o">CH2O: {envdata.ch2o} </p> }
      </div>
    </>
  )
}

export default Env
