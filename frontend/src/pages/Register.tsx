/* eslint-disable @typescript-eslint/no-explicit-any */
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";
import Button from "../components/Button";
import Input from "../components/Input";

const Register = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const handleRegister = async () => {
    if (!username || !password) {
      alert("Please fill in all fields");
      return;
    }

    setLoading(true);
    try {
      await axios.post("http://127.0.0.1:8081/register", {
        username,
        password,
      });
      alert("Registration successful. Please log in.");
      navigate("/");
    } catch (error: any) {
      alert(error.response?.data?.message || "Registration failed");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="container">
      <h2>Register</h2>
      <p>Sign up to continue</p>
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
      <Button
        text={loading ? "Registering..." : "Register"}
        onClick={handleRegister}
      />
      <p>
        Already have an account?{" "}
        <strong className="link">
          <span onClick={() => navigate("/")} className="clickable">
            Login
          </span>
        </strong>
      </p>
    </div>
  );
};

export default Register;
