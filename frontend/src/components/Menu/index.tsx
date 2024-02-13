import { View, Button } from "react-native";

export default function Menu({ navigation }: any) {
    return (
        <View>
            <Button
                title="Login"
                onPress={() =>
                    navigation.navigate("Login")
                }
            />
            <Button
                title="Register"
                onPress={() =>
                    navigation.navigate("Register")
                }
            />
        </View>
    );
};
