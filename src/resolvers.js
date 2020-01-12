require("dotenv").config(); //dotenv should be topmost so it loads all  env data
import Axios from "axios";
import bcrypt from "bcryptjs";
import jwt from "jsonwebtoken";
import { Storage } from "@google-cloud/storage";
import passport from "passport";
var GoogleStrategy = require("passport-google-oauth").OAuth2Strategy;
// import GoogleStrategy from "passport-google-oauth";

passport.use(
  new GoogleStrategy(
    {
      clientID: process.env.CLIENT_ID,
      clientSecret: process.env.CLIENT_SECRET,
      callbackURL: "http://www.example.com/auth/google/callback"
    },
    function(accessToken, refreshToken, profile, done) {
      User.findOrCreate({ googleId: profile.id }, function(err, user) {
        return done(err, user);
      });
    }
  )
);

const resolver = {
  Query: {
    info: () => `This is the API of a Hackernews Clone`,

    event: (_, ctx, prisma, info) => {
      const email = ctx.where.email;
      const id = ctx.where.id;
      return prisma.db.query.event({
        where: {
          email
        },
        info
      });
    },

    user: (_, ctx, prisma, info) => {
      const id = ctx.where.id;
      return prisma.db.query.user({
        where: {
          id
        },
        info
      });
    }
  },

  Mutation: {
    // ===================>
    createEvent: async (root, args, context, info) => {
      const hashedPassword = await bcrypt.hash(args.password, 10);

      // cloud function here
      const Email = args.email;
      try {
        Axios.post("http://localhost:8080/", {
          email: Email
        });
      } catch (error) {
        console.log(error);
      }

      return await context.db.mutation.createEvent({
        data: {
          name: args.name,
          description: args.description,
          // createdAt: new Date(),
          // use moment.js for createdAt
          type: args.type,
          venue: args.venue,
          duration: 11,
          organizer: args.organizer,
          website: args.website,
          bucketLink: args.bucketLink,
          supportEmail: args.supportEmail,
          teams: args.teams,
          attendees: args.attendees,
          password: hashedPassword
        }
      });
    },

    createUser: async (_, args, context, info) => {
      const hashedPassword = await bcrypt.hash(args.password, 10);

      // cloud function here
      const Email = args.email;
      try {
        Axios.post("http://localhost:8080/", {
          email: Email
        });
      } catch (error) {
        console.log(error);
      }

      return context.db.mutation.createStaff({
        data: {
          firstname: args.firstname,
          lastname: args.lastname,
          role: args.role,
          isLead: args.isLead,
          email: args.email,
          country: args.country,
          state: args.state,
          password: hashedPassword
        }
      });
    },

    //
    //  AUTH RESOLVERS ========>
    //

    loginUser: async (_, { password, where }, ctx, info) => {
      const email = where.email;
      const user = await ctx.db.query.staff({
        where: {
          email: email
        }
      });

      if (!user) {
        throw new Error("Invalid Login");
      }
      const passwordMatch = await bcrypt.compare(password, user.password);

      if (!passwordMatch) {
        throw new Error("Invalid Login");
      }
      const token = jwt.sign(
        {
          username: email
        },
        process.env.APP_SECRET,
        {
          expiresIn: "30d"
        }
      );
      return { token, user };
    },

    createGoogleUser: async (_, { email, where }, ctx, info) => {
      try {
        passport.authenticate("google", { scope: ["profile"] });

        console.log(passport);
      } catch (e) {
        console.log(e);
      }
      console.log(email);
      // return { user };
      return email;
    },

    // Cloud Functions resolvers here ============>>>>
    sendEmail: (root, args, context, info) => {
      const Email = args.email;
      try {
        Axios.post("http://localhost:8080/", {
          email: Email
        });
      } catch (error) {
        console.log(error);
      }
    }
  }
};

export { resolver };
