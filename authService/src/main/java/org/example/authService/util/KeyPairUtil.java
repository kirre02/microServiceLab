package org.example.authService.util;

import java.security.KeyPair;
import java.security.KeyPairGenerator;

public class KeyPairUtil {

    public static KeyPair generateRsaKey() {
        try {
            var generator = KeyPairGenerator.getInstance("RSA");
            generator.initialize(2048);
            return generator.generateKeyPair();
        } catch (Exception e) {
            throw new IllegalStateException("Could not generate RSA key", e);
        }
    }
}

