"use client";

import React, { useState } from "react";
import { Icon } from "@iconify/react";
import { Button } from "@/components/ui/button";
import deleteUser from './deleUser';
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { users } from "./data";
import { Checkbox } from "@/components/ui/checkbox";
import useFetchUsers from './userlist';
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Badge } from "@/components/ui/badge";
import {TableFooter} from "@mui/material";
import {getRoleFromJwt} from "./jwtUtils";
import getUser from "./getUser";
import EditUserForm from "../edit-user/validating-multiple-rules";

const CheckboxWithAction = () => {
  const [selectedRows, setSelectedRows] = useState([]);
  const { userList, loading } = useFetchUsers();
  const role = getRoleFromJwt();


  const handleSelectAll = (event) => {
    if (selectedRows?.length === userList?.length) {
      setSelectedRows([]);
    } else {
      setSelectedRows(userList.map((row) => row.ID));
    }
  };

  const handleRowSelect = (ID) => {
    const updatedSelectedRows = [...selectedRows];
    if (selectedRows.includes(ID)) {
      updatedSelectedRows.splice(selectedRows.indexOf(ID), 1);
    } else {
      updatedSelectedRows.push(ID);
    }
    setSelectedRows(updatedSelectedRows);
  };

  return (
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead>
            <Checkbox
              checked={selectedRows.length === userList.length || "indeterminate"}
              onCheckedChange={handleSelectAll}
            />
          </TableHead>

          <TableHead className=" font-semibold">
            {selectedRows.length === userList.length ? (
              <div className=" flex gap-2">
                <Button
                  size="xs"
                  variant="outline"
                  className=" text-xs "
                  color="secondary"
                >
                  Bulk edit
                </Button>
                <Button
                  size="xs"
                  variant="outline"
                  className=" text-xs "
                  color="destructive"
                >
                  Delete all
                </Button>
              </div>
            ) : (
              "userList"
            )}
          </TableHead>
          <TableHead> Title</TableHead>
          <TableHead>Email</TableHead>
          <TableHead>Role</TableHead>
          <TableHead>Action</TableHead>
        </TableRow>
      </TableHeader>

      <TableBody>
        {userList.map((item) => (
          <TableRow
            key={item.Email}
            className="hover:bg-muted"
            data-state={selectedRows.includes(item.ID) && "selected"}
          >
            <TableCell>
              <Checkbox
                checked={selectedRows.includes(item.ID)}
                onCheckedChange={() => handleRowSelect(item.ID)}
              />
            </TableCell>
            <TableCell className="  font-medium  text-card-foreground/80">
              <div className="flex gap-3 items-center">
                <Avatar className=" rounded-full">
                  <AvatarImage src={item.avatar} />
                  <AvatarFallback>AB</AvatarFallback>
                </Avatar>
                <span className=" text-sm   text-card-foreground">
                  {item.Name}
                </span>
              </div>
            </TableCell>

            <TableCell>{item.Job}</TableCell>
            <TableCell>{item.Email}</TableCell>
            <TableCell>
              <Badge
                variant="soft"
                color={
                  (item.Role === "admin" && "default") ||
                  (item.Role === "member" && "success") ||
                  (item.Role === "owner" && "info") ||
                  (item.Role === "editor" && "warning")
                }
                className=" capitalize"
              >
                {item.Role}
              </Badge>
            </TableCell>

            <TableCell className="flex justify-end">
              <div className="flex gap-3">
                {role === 'superadmin' && (
                    <Button
                        size="icon"
                        variant="outline"
                        color="secondary"
                        className="h-7 w-7"

                    >
                      <a href={`/edit-user`}> <Icon icon="heroicons:pencil" className="h-4 w-4"/></a>

                    </Button>
                )}
                {role === 'superadmin' && (
                    <Button
                        size="icon"
                        variant="outline"
                        className=" h-7 w-7"
                        color="secondary"
                        onClick={() => deleteUser(item.ID)}
                    >
                      <Icon icon="heroicons:trash" className=" h-4 w-4" />
                    </Button>
                )}



              </div>
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>

  );
};

export default CheckboxWithAction;
