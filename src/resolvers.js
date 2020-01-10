require("dotenv").config(); //dotenv should be topmost so it loads all  env data
import Axios from "axios";
import bcrypt from "bcryptjs";
import jwt from "jsonwebtoken";

const resolver = {
  // Subscription: {
  //   eventMessageSubscribe: {
  //     subscribe: (parent, args, ctx, info) => {
  //       console.log(ctx, "ctx");
  //       console.log(info, "info");
  //       return ctx.db.subscription.eventMessageSubscribe(
  //         {
  //           where: {
  //             mutation_in: ["CREATED", "UPDATED"]
  //           }
  //         },
  //         info
  //       );
  //     }
  //   }
  // },

  Query: {
    info: () => `This is the API of a Hackernews Clone`,

    // organization: async (_, ctx, prisma, info) => {
    // //   // logic to  create protected field ===========-->>
    // //
    // //   // if (!name) {
    // //   //   throw new Error('Invalid Login');
    // //   // }
    // //   const id = ctx.where.id;
    // //   const email = ctx.where.email;
    // //   return prisma.db.query.organization({
    // //     where: {
    // //       email
    // //     },
    // //     info
    // //   });
    // // },

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
          email: args.email,
          country: args.country,
          state: args.state,
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

    // updateGroup: async (_, args, context, info) => {
    //   const email = args.where.email;
    //   console.log(args.where.email);
    //
    //   return context.db.mutation.updateGroup({
    //     where: {
    //       email
    //     },
    //     data: {
    //       name: args.where.name,
    //       description: args.where.description,
    //       email: args.where.email,
    //       website: args.where.website,
    //       teams: args.where.teams,
    //       leads: args.where.leads
    //     }
    //   });
    // },

    //
    //  AUTH RESOLVERS ========>
    //

    loginEvent: async (_, { password, where }, ctx, info) => {
      const email = where.email;
      const user = await ctx.db.query.organization({
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
