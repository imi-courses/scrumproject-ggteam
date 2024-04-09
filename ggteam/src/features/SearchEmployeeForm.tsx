import { FormEvent, useEffect, useState } from "react";
import Button from "ui/Button";
import Input from "ui/Input";

const SearchEmployeeForm = () => {
 const [email, setEmail] = useState("");
 const [isAuth, setAuth] = useState(false);

 const onSubmit = async (e: FormEvent) => {
    e.preventDefault();


    const response = await fetch(
      import.meta.env.VITE_API_URL + "/employee/find_one",
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: "Bearer " + localStorage.getItem("access_token"),
        },
        credentials: "include",
      },
    );

    const json = await response.json();
    console.log(json);
 };

 useEffect(() => {
    if (window && localStorage.getItem("access_token")) {
      setAuth(true);
    }
 }, []);

 if (!isAuth) return;

 return (
    <form onSubmit={onSubmit}>
      <div>
        <span>Search by Email</span>
        <div>
          <Input value={email} setValue={setEmail} />
        </div>
      </div>
      <Button type="submit">Search</Button>
    </form>
 );
};

export default SearchEmployeeForm;