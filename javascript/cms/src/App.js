import { ColorModeContext, useMode } from "./theme";
import { CssBaseline, ThemeProvider } from "@mui/material";
import SideMenu from "./pages/global/SideMenu";

function App() {
  const [theme, colorMode] = useMode();

  return (
    <ColorModeContext.Provider value={colorMode}>
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <div className="app">
          <SideMenu />
          <main className="content">test</main>
        </div>
      </ThemeProvider>
    </ColorModeContext.Provider>
  );
}

export default App;
