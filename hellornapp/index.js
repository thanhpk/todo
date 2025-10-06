/**
 * @format
 */

import { AppRegistry } from 'react-native';
// import App from './App';
import 'react-native-gesture-handler';
import { name as appName } from './app.json';
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
// import { createStackNavigator } from '@react-navigation/stack';
const Stack = createNativeStackNavigator();
import { Suspense, lazy } from 'react';
const HomeScreen = lazy(() => import('./HomeScreen'));
const ProfileScreen = lazy(() => import('./ProfileScreen'));
import Loading from './Loading';

function MyStack() {
  return (
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
  );
}

AppRegistry.registerComponent(appName, () => MyStack);
