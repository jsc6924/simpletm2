import { Box, Flex, Heading, Spacer, Button, Link } from "@chakra-ui/react";
import { Link as RouterLink } from "react-router-dom";

const Header = () => {
  return (
    <Box bg="blue.500" color="white" px={4} py={2}>
      <Flex align="center">
        <Heading as="h1" size="lg">
          <Link as={RouterLink} to="/home">
            SimpleTM
          </Link>
        </Heading>
        <Spacer />
        <Link
          as={RouterLink}
          to="/"
          colorScheme="blue"
          variant="outline"
          ml={4}
        >
          Index
        </Link>
        <Button
          as={RouterLink}
          to="/login"
          variant="outline"
          ml={4}
          color="black"
        >
          Login
        </Button>
      </Flex>
    </Box>
  );
};

export default Header;
