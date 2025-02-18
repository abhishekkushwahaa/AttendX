import { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";
import Button from "../components/Button";
import Input from "../components/Input";

const Login = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleLogin = async () => {
    try {
      const response = await axios.post("http://127.0.0.1:8081/login", {
        username,
        password,
      });
      localStorage.setItem("authToken", response.data.token);
      navigate("/attendance");
    } catch {
      alert("Login failed");
    }
  };

  return (
    <div className="container">
      <h2>Welcome Back</h2>
      <p>Sign in to continue</p>
      <div className="input-container">
        <Input
          type="text"
          placeholder="Username"
          value={username}
          onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
            setUsername(e.target.value)
          }
        />
        <Input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
            setPassword(e.target.value)
          }
        />
      </div>
      <Button text="Login" onClick={handleLogin} />
      <p>
        Already have an account?{" "}
        <strong className="link">
          <span onClick={() => navigate("/register")} className="clickable">
            Register
          </span>
        </strong>
      </p>
    </div>
  );
};

export default Login;
