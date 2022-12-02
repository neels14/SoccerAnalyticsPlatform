import { useState, useContext } from "react";
import Box from "@mui/material/Box";
import { TextField } from "@mui/material";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import UserContext from "./shared/user-context";
import { Link, Navigate, useMatch } from "react-router-dom";

function Login() {
    const [errorUsername, setErrorUsername] = useState(false)
    const [errorPassword, setErrorPassword] = useState(false)
    const { user, setUser } = useContext(UserContext)
    const isLoginURL = useMatch({ path: `/login`, end: true})

    const handleSubmit = (event) => {
        setErrorUsername(false)
        setErrorPassword(false)

        event.preventDefault();

        const data = new FormData(event.currentTarget);
        const username = data.get('username')
        const password = data.get('password')
        if (username === "") {
            setErrorUsername(true)
        }
        if (password === "") {
            setErrorPassword(true)
        }
        if (username === "" || password === "") {
            return
        }

        fetch(`/api/v1/user/${username}/${password}`)
            .then(response => response.json())
            .then(json => {
                if (json.valid) {
                    setUser(username)
                } else {
                    setErrorUsername(true)
                    setErrorPassword(true)
                }
            })
    }

    return (
        <>
        {user === "" ?
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
                {!isLoginURL &&
                    <Typography variant="h4">Please Login to Continue</Typography>
                }
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
                <Button type="submit" size="large" variant="outlined"><b>Login</b></Button>
                <Typography>
                Don't have an account?<Button component={Link} to={`/signup`}>Sign Up</Button>
                </Typography>
            </Box> :
            <Navigate to="/" />
        }
        </>
    );
  }
  
export default Login;
