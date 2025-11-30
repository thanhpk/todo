import React from 'react';
import { createRoot } from 'react-dom/client';
import { html, css } from 'react-strict-dom';
import 'stylex-webpack/stylex.css';
import './root.css';

import 'react-native-gesture-handler';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
const Stack = createStackNavigator();
import { Suspense, lazy } from 'react';
const HomeScreen = () => <div>thanh</div>;
const ProfileScreen = () => <div>profile</div>;
const Loading = () => <div>loading</div>;

const styles = css.create({
  root: {
    backgroundColor: 'lightblue',
    padding: '20px',
  },
  text: {
    color: 'red',
    fontSize: '16px',
  },
  container: {
    display: 'flex',
  },
  child1style: {
    backgroundColor: 'yellow',
    width: '100px',
    height: '100px',
    color: 'green',
  },
  child2style: {
    backgroundColor: 'white',
    width: '100px',
    height: '100px',
    color: 'black',
    fontSize: '12px',
  },
});

function App() {
  let name = 'thanh';
  return (
    <html.div style={styles.root}>
      <html.p style={styles.text}>Hello, React Strict DOM!{name}</html.p>
      <html.p>Text 2</html.p>
      <html.div style={styles.container}>
        <html.p style={styles.child1style}>Child 1</html.p>
        <html.p style={styles.child2style}>Child 2</html.p>
      </html.div>
    </html.div>
  );
}

const container = document.getElementById('root');
const root = createRoot(container);

root.render(
  <React.StrictMode>
    <NavigationContainer>
      <Stack.Navigator
        screenLayout={({ children }) => (
          <Suspense fallback={<Loading />}>{children}</Suspense>
        )}
      >
        <Stack.Screen name="Home" component={HomeScreen} />
        <Stack.Screen name="Profile" component={ProfileScreen} />
      </Stack.Navigator>
    </NavigationContainer>
  </React.StrictMode>,
);
