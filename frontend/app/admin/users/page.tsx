"use client";

import CustomButton from "@/app/components/button";
import { Dropdown } from "@/app/components/dropdown";
import { InputWithLabel } from "@/app/components/input";

const users = [
  { id: 1, name: "Shuttlers IT", email: "it@shuttlers.ng" },
  { id: 2, name: "Shuttlers IT", email: "it@shuttlers.ng" },
];

const Users = () => {
  return (
    <div className="shadow-[0px_0px_32px_0px_rgba(204,204,204,0.25)] rounded-2xl px-10 py-10">
      <form
        action=""
        className="flex gap-5 border-b border-[#EBEBEB] pb-5 mb-5"
      >
        <InputWithLabel className="rounded-2xl" label="Name" />
        <InputWithLabel className="rounded-2xl" label="Email" />
      </form>
      <div className="flex flex-col gap-5 border-b border-[#EBEBEB] pb-5 mb-5">
        {users.map((user) => (
          <div key={user.id} className="flex justify-between items-center">
            <p>{user.name}</p>
            <p>{user.email}</p>
            <div className="flex gap-5">
              <CustomButton className="bg-[#0AAE5B] hover:bg-[#7ee1b0]">
                Edit
              </CustomButton>
              <CustomButton>Reset Password</CustomButton>
            </div>
          </div>
        ))}
      </div>
      <div className="flex justify-between">
        <Dropdown />
        <div className="gap-5 flex">
          <CustomButton>Previous</CustomButton>
          <CustomButton>Next</CustomButton>
        </div>
      </div>
    </div>
  );
};

export default Users;
