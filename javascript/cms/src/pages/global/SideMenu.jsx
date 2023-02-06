import { useState } from "react";
import { Link } from "react-router-dom";
import { Box, Select, useTheme, IconButton, Typography } from "@mui/material";
import MaterialMenuItem from "@mui/material/MenuItem";
import { tokens } from "../../theme";
import { Sidebar, Menu, MenuItem } from "react-pro-sidebar";
import MenuOutlinedIcon from "@mui/icons-material/MenuOutlined";
import SchemaOutlinedIcon from "@mui/icons-material/SchemaOutlined";
import WebAssetOutlinedIcon from "@mui/icons-material/WebAssetOutlined";
import ContentPasteOutlinedIcon from "@mui/icons-material/ContentPasteOutlined";
import ClassOutlinedIcon from "@mui/icons-material/ClassOutlined";

const Item = ({ title, to, icon, selected, setSelected }) => {
  const theme = useTheme();
  const colors = tokens(theme.palette.mode);
  return (
    <MenuItem
      component={<Link to={to} />}
      active={selected === title}
      style={{
        color: colors.grey[100],
      }}
      onClick={() => setSelected(title)}
      icon={icon}
    >
      <Typography>{title}</Typography>
    </MenuItem>
  );
};

export default function SideMenu() {
  const theme = useTheme();
  const colors = tokens(theme.palette.mode);
  const [isCollapsed, setIsCollapsed] = useState(false);
  const [selected, setSelected] = useState("Projects");

  return (
    <Box
      sx={{
        "& .ps-sidebar-root": {
          borderColor: `transparent !important`,
        },
        "& .ps-sidebar-container": {
          background: `${colors.primary[400]} !important`,
        },
        "& .pro-icon-wrapper": {
          backgroundColor: "transparent !important",
        },
        "& .pro-inner-item": {
          padding: "5px 35px 5px 20px !important",
        },
        "& .ps-menu-button:hover": {
          color: "#868dfb !important",
          backgroundColor: "transparent !important",
        },
        "& .pro-menu-item.active": {
          color: "#6870fa !important",
        },
      }}
    >
      <Sidebar collapsed={isCollapsed}>
        <Menu>
          <MenuItem
            onClick={() => setIsCollapsed(!isCollapsed)}
            icon={isCollapsed ? <MenuOutlinedIcon /> : undefined}
            style={{
              margin: "10px 0 20px 0",
              color: colors.grey[100],
            }}
          >
            {!isCollapsed && (
              <Box
                display="flex"
                justifyContent="space-between"
                alignItems="center"
                ml="15px"
              >
                <Typography variant="h3" color={colors.grey[100]}>
                  MY CMS
                </Typography>
                <IconButton onClick={() => setIsCollapsed(!isCollapsed)}>
                  <MenuOutlinedIcon />
                </IconButton>
              </Box>
            )}
          </MenuItem>

          <Typography
            variant="h6"
            color={colors.grey[300]}
            sx={{ m: "15px 0 5px 20px" }}
          >
            Project
          </Typography>

          <Box paddingLeft={isCollapsed ? undefined : "10%"}>
            <MenuItem>
              <Select
                labelId="demo-simple-select-label"
                id="demo-simple-select"
                label="Project"
              >
                <MaterialMenuItem value={10}>Ten</MaterialMenuItem>
                <MaterialMenuItem value={20}>Twenty</MaterialMenuItem>
                <MaterialMenuItem value={30}>Thirty</MaterialMenuItem>
              </Select>
            </MenuItem>
          </Box>

          <Box paddingLeft={isCollapsed ? undefined : "10%"}>
            <Item
              title="Projects"
              to="/"
              icon={<ClassOutlinedIcon />}
              selected={selected}
              setSelected={setSelected}
            />
          </Box>

          <Typography
            variant="h6"
            color={colors.grey[300]}
            sx={{ m: "15px 0 5px 20px" }}
          >
            Management
          </Typography>

          <Box paddingLeft={isCollapsed ? undefined : "10%"}>
            <Item
              title="Schema"
              to="/"
              icon={<SchemaOutlinedIcon />}
              selected={selected}
              setSelected={setSelected}
            />
          </Box>

          <Box paddingLeft={isCollapsed ? undefined : "10%"}>
            <Item
              title="Content"
              to="/"
              icon={<ContentPasteOutlinedIcon />}
              selected={selected}
              setSelected={setSelected}
            />
          </Box>

          <Box paddingLeft={isCollapsed ? undefined : "10%"}>
            <Item
              title="Assets"
              to="/"
              icon={<WebAssetOutlinedIcon />}
              selected={selected}
              setSelected={setSelected}
            />
          </Box>
        </Menu>
      </Sidebar>
    </Box>
  );
}
