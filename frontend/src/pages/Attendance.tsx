import axios from "axios";
import Button from "../components/Button";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

const Attendance = () => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("authToken");
    if (!token) {
      navigate("/");
    } else {
      setIsAuthenticated(true);
    }
  }, [navigate]);
  const handleFingerprintAuth = async () => {
    try {
      const options = await axios.post(
        "http://127.0.0.1:8081/auth/fingerprint"
      );
      const publicKeyCredential = await navigator.credentials.create({
        publicKey: options.data,
      });

      await axios.post("http://127.0.0.1:8081/auth/fingerprint/verify", {
        credential: publicKeyCredential,
      });
      alert("Attendance marked successfully");
    } catch {
      alert("Fingerprint authentication failed");
    }
  };

  if (!isAuthenticated) {
    return null;
  }

  return (
    <div className="container">
      <h2>Mark Attendance</h2>
      <Button text="Verify Fingerprint" onClick={handleFingerprintAuth} />
    </div>
  );
};

export default Attendance;
