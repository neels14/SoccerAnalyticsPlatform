import { useState } from "react";
import Box from "@mui/material/Box";
import { TextField } from "@mui/material";
import Button from "@mui/material/Button";
import { useNavigate } from "react-router-dom";

function SignUp() {
    const [errorUsername, setErrorUsername] = useState(false)
    const [errorPassword, setErrorPassword] = useState(false)
    const [errorFirstName, setErrorFirstName] = useState(false)
    const [errorLastName, setErrorLastName] = useState(false)
    const [errorEmail, setErrorEmail] = useState(false)
    const navigate = useNavigate();

    const handleSubmit = (event) => {
        setErrorFirstName(false)
        setErrorLastName(false)
        setErrorEmail(false)
        setErrorUsername(false)
        setErrorPassword(false)

        event.preventDefault();

        const data = new FormData(event.currentTarget);
        const firstname = data.get('first-name')
        const lastname = data.get('last-name')
        const email = data.get('email')
        const username = data.get('username')
        const password = data.get('password')
        if (firstname === "") {
            setErrorFirstName(true)
        }
        if (lastname === "") {
            setErrorLastName(true)
        }
        if (email === "") {
            setErrorEmail(true)
        }
        if (username === "") {
            setErrorUsername(true)
        }
        if (password === "") {
            setErrorPassword(true)
        }
        if (firstname === "" || lastname === "" || email === "" ||username === "" || password === "") {
            return
        }

        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ 
                username: username,
                first_name: firstname,
                last_name: lastname,
                email: email,
                password: password 
            })
        };

        fetch(`/api/v1/user/`, requestOptions)
            .then(response => response.json())
            .then(json => {
                if (json === null) {
                    setErrorEmail(true)
                    setErrorUsername(true)
                } else {
                    navigate("/login")
                }
            })
    }

    return (
        <Box
            component="form"
            sx={{
                '& .MuiTextField-root': { m: 2, width: '30ch' },
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
                mt: "50px",
                width: 1
            }}
            noValidate
            autoComplete="off"
            onSubmit={handleSubmit}
        >
            <TextField
                required
                id="first-name"
                label="First Name"
                name="first-name"
                error={errorFirstName}
            />
            <TextField
                required
                id="last-name"
                label="Last Name"
                name="last-name"
                error={errorLastName}
            />
            <TextField
                required
                id="email"
                label="Email"
                name="email"
                error={errorEmail}
            />
            <TextField
                required
                id="username"
                label="Username"
                name="username"
                error={errorUsername}
            />
            <TextField
                required
                id="password"
                label="Password"
                name="password"
                type="password"
                error={errorPassword}
            />
            <Button type="submit" size="large" variant="outlined"><b>Sign Up</b></Button>
        </Box>
    );
  }
  
export default SignUp;
